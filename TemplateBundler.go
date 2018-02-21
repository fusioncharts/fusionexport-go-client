package FusionExport

import (
	"archive/zip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type TemplateBundler struct {
	Template           string
	Resources          string
	templatePathInZip  string
	resourcesZipFile   string
	basePath           string
	parsedResources    resources
	collectedResources []string
	zipMetas           []zipMeta
}

type resources struct {
	BasePath string   `json:"basePath"`
	Include  []string `json:"include"`
	Exclude  []string `json:"exclude"`
}

type zipMeta struct {
	path    string
	zipPath string
}

func resolveGlobs(paths []string, basePath string) ([]string, error) {
	var jps []string
	var rps []string
	var aps []string

	if len(paths) == 0 {
		return aps, nil
	}

	for _, v := range paths {
		jv := filepath.Join(basePath, v)
		jps = append(jps, jv)
	}

	for _, v := range jps {
		gps, err := filepath.Glob(v)
		if err != nil {
			return nil, err
		}

		rps = append(rps, gps...)
	}

	for _, v := range rps {
		ap, err := filepath.Abs(v)
		if err != nil {
			return nil, err
		}

		aps = append(aps, ap)
	}

	return aps, nil
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func findCommonPath(paths []string) string {
	sort.Strings(paths)
	if len(paths) == 0 {
		return ""
	}
	p1 := strings.Split(paths[0], string(os.PathSeparator))
	p2 := strings.Split(paths[len(paths)-1], string(os.PathSeparator))
	l := len(p1)
	i := 0
	for i < l && p1[i] == p2[i] {
		i++
	}
	return strings.Join(p1[:i], string(os.PathSeparator))
}

func removeCommonPath(base, common string) string {
	baseSplit := strings.Split(base, string(os.PathSeparator))
	commonSplit := strings.Split(common, string(os.PathSeparator))
	l := len(commonSplit)
	i := 0
	for i < l && i < len(baseSplit) && baseSplit[i] == commonSplit[i] {
		i++
	}
	return strings.Join(baseSplit[i:], string(os.PathSeparator))
}

func (tb *TemplateBundler) Process() error {
	if tb.Template == "" {
		return errors.New("no template found")
	}

	err := tb.findHTMLResources()
	if err != nil {
		return err
	}

	err = tb.resolveResourcesGlobs()
	if err != nil {
		return err
	}

	err = tb.zipResourcesFiles()
	if err != nil {
		return err
	}

	return nil
}

func (tb *TemplateBundler) GetTemplatePathInZip() string {
	return tb.templatePathInZip
}

func (tb *TemplateBundler) GetResourcesZip() string {
	return tb.resourcesZipFile
}

func (tb *TemplateBundler) GetResourcesZipAsBase64() (string, error) {
	data, err := ioutil.ReadFile(tb.resourcesZipFile)
	if err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString([]byte(data))
	return b64, nil
}

func (tb *TemplateBundler) findHTMLResources() error {
	data, err := ioutil.ReadFile(tb.Template)
	if err != nil {
		return err
	}

	htmlParser := HTMLParser{Data: data}

	links, err := htmlParser.GetElemetsByTagName("link")
	if err != nil {
		return err
	}

	scripts, err := htmlParser.GetElemetsByTagName("script")
	if err != nil {
		return err
	}

	imgs, err := htmlParser.GetElemetsByTagName("img")
	if err != nil {
		return err
	}

	for _, link := range links {
		tb.collectedResources = append(tb.collectedResources, link.GetAttribute("href"))
	}

	for _, script := range scripts {
		tb.collectedResources = append(tb.collectedResources, script.GetAttribute("src"))
	}

	for _, img := range imgs {
		tb.collectedResources = append(tb.collectedResources, img.GetAttribute("src"))
	}

	tb.filterRemoteCollectedResources()

	err = tb.resolveCollectedResources()
	if err != nil {
		return err
	}

	tb.filterDuplicateCollectedResources()

	return nil
}

func (tb *TemplateBundler) filterRemoteCollectedResources() {
	var filteredResources []string

	for _, v := range tb.collectedResources {
		if v == "" {
			continue
		}

		if strings.HasPrefix(v, "http://") {
			continue
		}

		if strings.HasPrefix(v, "https://") {
			continue
		}

		if strings.HasPrefix(v, "file://") {
			continue
		}

		filteredResources = append(filteredResources, v)
	}

	tb.collectedResources = filteredResources
}

func (tb *TemplateBundler) resolveCollectedResources() error {
	var resolvedResources []string

	tmplDir := filepath.Dir(tb.Template)
	tmplAbsDir, err := filepath.Abs(tmplDir)
	if err != nil {
		return err
	}

	for _, v := range tb.collectedResources {
		if filepath.IsAbs(v) {
			resolvedResources = append(resolvedResources, v)
		} else {
			jv := filepath.Join(tmplAbsDir, v)
			resolvedResources = append(resolvedResources, jv)
		}
	}

	tb.collectedResources = resolvedResources
	return nil
}

func (tb *TemplateBundler) filterDuplicateCollectedResources() {
	var uniques []string

	keys := make(map[string]bool)
	for _, entry := range tb.collectedResources {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniques = append(uniques, entry)
		}
	}

	tb.collectedResources = uniques
}

func (tb *TemplateBundler) resolveResourcesGlobs() error {
	err := tb.parseResourcesFile()
	if err != nil {
		return err
	}

	if &tb.parsedResources == nil {
		return nil
	}

	resourcesDir := filepath.Dir(tb.Resources)

	includeFiles, err := resolveGlobs(tb.parsedResources.Include, resourcesDir)
	if err != nil {
		return nil
	}

	excludeFiles, err := resolveGlobs(tb.parsedResources.Exclude, resourcesDir)
	if err != nil {
		return nil
	}

	var resourcesFiles []string
	for _, v := range includeFiles {
		if !stringInSlice(v, excludeFiles) {
			resourcesFiles = append(resourcesFiles, v)
		}
	}

	tb.collectedResources = append(tb.collectedResources, resourcesFiles...)

	return nil
}

func (tb *TemplateBundler) parseResourcesFile() error {
	if tb.Resources == "" {
		return nil
	}

	resourcesBytes, err := ioutil.ReadFile(tb.Resources)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resourcesBytes, &tb.parsedResources)
	if err != nil {
		return err
	}

	return nil
}

func (tb *TemplateBundler) zipResourcesFiles() error {
	err := tb.sanitizeBasePath()
	if err != nil {
		return nil
	}

	var filteredResources []string
	for _, v := range tb.collectedResources {
		if strings.HasPrefix(v, tb.basePath) {
			filteredResources = append(filteredResources, v)
		}
	}

	var zms []zipMeta
	for _, v := range filteredResources {
		zm := zipMeta{
			path:    v,
			zipPath: removeCommonPath(v, tb.basePath),
		}
		zms = append(zms, zm)
	}
	tb.zipMetas = zms

	err = tb.insertTemplateInZipMetas()
	if err != nil {
		return err
	}

	zipFile, err := ioutil.TempFile("", "fe-")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, v := range tb.zipMetas {
		f, err := os.Open(v.path)
		if err != nil {
			return err
		}

		info, err := f.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, f)
		if err != nil {
			return err
		}

		f.Close()
	}

	tb.resourcesZipFile = zipFile.Name()
	return nil
}

func (tb *TemplateBundler) sanitizeBasePath() error {
	if tb.parsedResources.BasePath != "" {
		tb.basePath = tb.parsedResources.BasePath
	} else {
		tb.basePath = findCommonPath(tb.collectedResources)
	}

	bp, err := filepath.Abs(tb.basePath)
	if err != nil {
		return err
	}

	tb.basePath = bp
	return nil
}

func (tb *TemplateBundler) insertTemplateInZipMetas() error {
	absTmplPath, err := filepath.Abs(tb.Template)
	if err != nil {
		return err
	}

	var zp string
	if len(tb.zipMetas) == 0 {
		zp = filepath.Base(absTmplPath)
	} else {
		zp = removeCommonPath(absTmplPath, tb.basePath)
	}

	zm := zipMeta{
		path:    absTmplPath,
		zipPath: zp,
	}

	tb.templatePathInZip = zp
	tb.zipMetas = append(tb.zipMetas, zm)
	return nil
}
