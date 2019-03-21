package search

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	client "github.com/guessi/kubectl-search/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	daemonsetsFields     = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE"
	daemonsetsFieldsWide = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	dsInfo               string
)

// Daemonsets - a public function for searching daemonsets with keyword
func Daemonsets(namespace string, allNamespaces bool, selector, fieldSelector, keyword string, wide bool) {
	clientset := client.InitClient()

	if len(namespace) <= 0 {
		namespace = "default"
	}

	if allNamespaces {
		namespace = ""
	}

	listOptions := &metav1.ListOptions{
		LabelSelector: selector,
		FieldSelector: fieldSelector,
	}

	daemonsets, err := clientset.AppsV1().DaemonSets(namespace).List(*listOptions)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, daemonsetsFieldsWide)
	} else {
		fmt.Fprintln(w, daemonsetsFields)
	}

	for _, d := range daemonsets.Items {
		// return all daemonsets under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(d.Name, keyword)
			if !match {
				continue
			}
		}

		age, ageUnit := getAge(time.Since(d.CreationTimestamp.Time).Seconds())
		containers := d.Spec.Template.Spec.Containers

		var nodeSelectors []string
		var nodeSelector string
		if d.Spec.Template.Spec.NodeSelector != nil {
			for k, v := range d.Spec.Template.Spec.NodeSelector {
				nodeSelector = fmt.Sprintf("%s=%s", k, v)
				nodeSelectors = append(nodeSelectors, nodeSelector)
			}
		}
		nodeSelectorOutput := "<none>"
		if len(nodeSelectors) > 0 {
			nodeSelectorOutput = strings.Join(nodeSelectors, ",")
		}

		var selectors []string
		var selector string
		if d.Spec.Selector.MatchLabels != nil {
			for k, v := range d.Spec.Selector.MatchLabels {
				selector = fmt.Sprintf("%s=%s", k, v)
				selectors = append(selectors, selector)
			}
		}
		selectorOutput := "<none>"
		if len(selectors) > 0 {
			selectorOutput = strings.Join(selectors, ",")
		}

		if wide {
			names := []string{}
			images := []string{}

			for _, n := range containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			dsInfo = fmt.Sprintf("%s\t%s\t%d\t%d\t%d\t%d\t%s\t%d%s\t%s\t%s\t%s",
				d.Namespace,
				d.Name,
				d.Status.DesiredNumberScheduled,
				d.Status.NumberReady,
				d.Status.UpdatedNumberScheduled,
				d.Status.NumberAvailable,
				nodeSelectorOutput,
				age, ageUnit,
				strings.Join(names, ","),
				strings.Join(images, ","),
				selectorOutput,
			)
		} else {
			dsInfo = fmt.Sprintf("%s\t%s\t%d\t%d\t%d\t%d\t%s\t%d%s",
				d.Namespace,
				d.Name,
				d.Status.DesiredNumberScheduled,
				d.Status.NumberReady,
				d.Status.UpdatedNumberScheduled,
				d.Status.NumberAvailable,
				nodeSelectorOutput,
				age, ageUnit,
			)
		}
		fmt.Fprintln(w, dsInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
