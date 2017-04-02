package job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	batchv1 "k8s.io/client-go/pkg/apis/batch/v1"
	"k8s.io/client-go/rest"
)

const jobTemplate = "/templates/job_resource.tmpl"

type Job struct {
	Name    string
	Image   string
	Command []string
}

func createJob(j *batchv1.Job) (err error) {
	c := k8s_client()
	_, err = c.BatchV1Client.Jobs("default").Create(j)
	return
}

func formatAsCommand(cmd string) (ret string) {
	for _, f := range strings.Fields(cmd) {
		ret += fmt.Sprintf("\"%s\",", f)
	}
	ret = "[" + strings.TrimRight(ret, ",") + "]"
	return ret
}

func getJobs() {
	c := k8s_client()
	jobs, err := c.BatchV1Client.Jobs("").List(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Jobs: %v", jobs)
}

func k8s_client() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func (j *Job) toK8SJob(tmpl string) (ret *batchv1.Job, err error) {
	ret = new(batchv1.Job)
	buf, err := render(j, tmpl)
	data, err := yaml.YAMLToJSON(buf.Bytes())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, ret)
	return ret, err
}

func render(j *Job, tmpl string) (buf bytes.Buffer, err error) {
	if tmpl == "" {
		tmpl = jobTemplate
	}
	fmap := template.FuncMap{"formatAsCommand": formatAsCommand}
	t := template.Must(template.New("job_resource.tmpl").Funcs(fmap).ParseFiles(tmpl))
	if err = t.Execute(&buf, j); err != nil {
		log.Println("Template Execute Error:", err)
		return buf, err
	}
	return buf, nil
}

func (j *Job) Run() error {
	log.Printf("Running job: %+v\n", j)
	job, err := j.toK8SJob("")
	if err != nil {
		panic(err.Error())
	}
	log.Printf("K8S Job:\n%#v\n", *job)
	err = createJob(job)
	if err != nil {
		log.Println("createJob() Error:", err)
	}
	return err
}
