package job

import (
	"testing"

	"github.com/ghodss/yaml"
)

func TestK8SJob(t *testing.T) {
	j := Job{Name: "test_job_name", Image: "test_image", Command: []string{"my", "command line"}}
	k8sJob, err := j.toK8SJob("../templates/job_resource.tmpl")
	if err != nil {
		t.Fail()
	}
	t.Logf("Kubernetes Job:\n%#v", k8sJob.Spec.Template.Spec.Containers[0].Command)
}

func TestRender(t *testing.T) {
	j := Job{Name: "test_job_name", Image: "test_image", Command: []string{"my", "command line"}}
	tmpl, err := render(&j, "../templates/job_resource.tmpl")
	if err != nil {
		t.Fail()
	}
	data, err := yaml.YAMLToJSON(tmpl.Bytes())
	t.Logf("Template:\n%s", tmpl.String())
	t.Logf("JSON:\n%s", string(data))
}
