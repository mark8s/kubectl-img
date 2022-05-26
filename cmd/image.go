/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/mark8s/kubectl-img/pkg/kube"
	"github.com/spf13/cobra"
	app "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:     "image",
	Short:   "kubernetes workload image",
	Long:    `.`,
	Example: "kubectl-img image -n all",
	Run: func(cmd *cobra.Command, args []string) {
		clientSet := kube.InitClientSet()

		namespace, _ := cmd.Flags().GetString("namespace")
		name, _ := cmd.Flags().GetString("name")
		kind, _ := cmd.Flags().GetString("type")
		all, _ := cmd.Flags().GetBool("all")

		var resources []interface{}

		switch kind {
		case "pod":
			if all {
				list, err := clientSet.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
				if err != nil {
					return
				}
				if len(list.Items) > 0 {
					for i := 0; i < len(list.Items); i++ {
						pods, err := clientSet.CoreV1().Pods(list.Items[i].Name).List(context.TODO(), v1.ListOptions{})
						if err != nil {
							return
						}
						resources = append(resources, pods)
					}
				}
			} else {
				if name == "" {
					pods, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
					if err != nil {
						return
					}
					resources = append(resources, pods)
				} else {
					pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, v1.GetOptions{})
					if err != nil {
						return
					}
					resources = append(resources, pod)
				}
			}
		case "deploy":
			if all {
				list, err := clientSet.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
				if err != nil {
					return
				}
				if len(list.Items) > 0 {
					for i := 0; i < len(list.Items); i++ {
						deploys, err := clientSet.AppsV1().Deployments(list.Items[i].Name).List(context.TODO(), v1.ListOptions{})
						if err != nil {
							return
						}
						resources = append(resources, deploys)
					}
				}
			} else {
				if name == "" {
					deploys, err := clientSet.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
					if err != nil {
						return
					}
					resources = append(resources, deploys)
				} else {
					deploy, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, v1.GetOptions{})
					if err != nil {
						return
					}
					resources = append(resources, deploy)
				}
			}
		case "sts":
			if all {
				list, err := clientSet.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
				if err != nil {
					return
				}
				if len(list.Items) > 0 {
					for i := 0; i < len(list.Items); i++ {
						sts, err := clientSet.AppsV1().StatefulSets(list.Items[i].Name).List(context.TODO(), v1.ListOptions{})
						if err != nil {
							return
						}
						resources = append(resources, sts)
					}
				}
			} else {
				if name == "" {
					sts, err := clientSet.AppsV1().StatefulSets(namespace).List(context.TODO(), v1.ListOptions{})
					if err != nil {
						return
					}
					resources = append(resources, sts)
				} else {
					sts, err := clientSet.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, v1.GetOptions{})
					if err != nil {
						return
					}
					resources = append(resources, sts)
				}
			}

		case "ds":
			if all {
				list, err := clientSet.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
				if err != nil {
					return
				}
				if len(list.Items) > 0 {
					for i := 0; i < len(list.Items); i++ {
						ds, err := clientSet.AppsV1().DaemonSets(list.Items[i].Name).List(context.TODO(), v1.ListOptions{})
						if err != nil {
							return
						}
						resources = append(resources, ds)
					}
				}
			} else {
				if name == "" {
					ds, err := clientSet.AppsV1().DaemonSets(namespace).List(context.TODO(), v1.ListOptions{})
					if err != nil {
						return
					}
					resources = append(resources, ds)
				} else {
					ds, err := clientSet.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, v1.GetOptions{})
					if err != nil {
						return
					}
					resources = append(resources, ds)
				}
			}
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return
		}
		var resourceMapList []map[string]string

		for i := 0; i < len(resources); i++ {
			switch t := resources[i].(type) {
			case *core.PodList:
				for j := 0; j < len(t.Items); j++ {
					resourceMap := make(map[string]string)
					resourceMap["NAMESPACE"] = t.Items[j].Namespace
					resourceMap["TYPE"] = "pod"
					resourceMap["NAME"] = t.Items[j].Name
					containers := t.Items[j].Spec.Containers
					var image string
					if len(containers) == 1 {
						image = containers[0].Image
					} else {
						for _, container := range containers {
							if container.Name == "istio-proxy" {
								continue
							}
							image = container.Image
							break
						}
					}
					resourceMap["IMAGE"] = image
					resourceMapList = append(resourceMapList, resourceMap)
				}
			case *core.Pod:
				resourceMap := make(map[string]string)
				resourceMap["NAMESPACE"] = t.Namespace
				resourceMap["TYPE"] = "pod"
				resourceMap["NAME"] = t.Name
				containers := t.Spec.Containers
				var image string
				if len(containers) == 1 {
					image = containers[0].Image
				} else {
					for _, container := range containers {
						if container.Name == "istio-proxy" {
							continue
						}
						image = container.Image
						break
					}
				}
				resourceMap["IMAGE"] = image
				resourceMapList = append(resourceMapList, resourceMap)

			case *app.Deployment:
				resourceMap := make(map[string]string)
				resourceMap["NAMESPACE"] = t.Namespace
				resourceMap["TYPE"] = "deploy"
				resourceMap["NAME"] = t.Name
				containers := t.Spec.Template.Spec.Containers
				var image string
				if len(containers) == 1 {
					image = containers[0].Image
				} else {
					for _, container := range containers {
						if container.Name == "istio-proxy" {
							continue
						}
						image = container.Image
						break
					}
				}
				resourceMap["IMAGE"] = image
				resourceMapList = append(resourceMapList, resourceMap)

			case *app.DeploymentList:
				for j := 0; j < len(t.Items); j++ {
					resourceMap := make(map[string]string)
					resourceMap["NAMESPACE"] = t.Items[j].Namespace
					resourceMap["TYPE"] = "deploy"
					resourceMap["NAME"] = t.Items[j].Name
					containers := t.Items[j].Spec.Template.Spec.Containers
					var image string
					if len(containers) == 1 {
						image = containers[0].Image
					} else {
						for _, container := range containers {
							if container.Name == "istio-proxy" {
								continue
							}
							image = container.Image
							break
						}
					}
					resourceMap["IMAGE"] = image
					resourceMapList = append(resourceMapList, resourceMap)
				}

			case *app.StatefulSet:
				resourceMap := make(map[string]string)
				resourceMap["NAMESPACE"] = t.Namespace
				resourceMap["TYPE"] = "sts"
				resourceMap["NAME"] = t.Name
				containers := t.Spec.Template.Spec.Containers
				var image string
				if len(containers) == 1 {
					image = containers[0].Image
				} else {
					for _, container := range containers {
						if container.Name == "istio-proxy" {
							continue
						}
						image = container.Image
						break
					}
				}
				resourceMap["IMAGE"] = image
				resourceMapList = append(resourceMapList, resourceMap)

			case *app.StatefulSetList:
				for j := 0; j < len(t.Items); j++ {
					resourceMap := make(map[string]string)
					resourceMap["NAMESPACE"] = t.Items[j].Namespace
					resourceMap["TYPE"] = "sts"
					resourceMap["NAME"] = t.Items[j].Name
					containers := t.Items[j].Spec.Template.Spec.Containers
					var image string
					if len(containers) == 1 {
						image = containers[0].Image
					} else {
						for _, container := range containers {
							if container.Name == "istio-proxy" {
								continue
							}
							image = container.Image
							break
						}
					}
					resourceMap["IMAGE"] = image
					resourceMapList = append(resourceMapList, resourceMap)
				}
			case *app.DaemonSet:
				resourceMap := make(map[string]string)
				resourceMap["NAMESPACE"] = t.Namespace
				resourceMap["TYPE"] = "ds"
				resourceMap["NAME"] = t.Name
				containers := t.Spec.Template.Spec.Containers
				var image string
				if len(containers) == 1 {
					image = containers[0].Image
				} else {
					for _, container := range containers {
						if container.Name == "istio-proxy" {
							continue
						}
						image = container.Image
						break
					}
				}
				resourceMap["IMAGE"] = image
				resourceMapList = append(resourceMapList, resourceMap)
			case *app.DaemonSetList:
				for j := 0; j < len(t.Items); j++ {
					resourceMap := make(map[string]string)
					resourceMap["NAMESPACE"] = t.Items[j].Namespace
					resourceMap["TYPE"] = "ds"
					resourceMap["NAME"] = t.Items[j].Name
					containers := t.Items[j].Spec.Template.Spec.Containers
					var image string
					if len(containers) == 1 {
						image = containers[0].Image
					} else {
						for _, container := range containers {
							if container.Name == "istio-proxy" {
								continue
							}
							image = container.Image
							break
						}
					}
					resourceMap["IMAGE"] = image
					resourceMapList = append(resourceMapList, resourceMap)
				}
			}
		}

		table := kube.GenTable(resourceMapList)
		if format == "table" {
			fmt.Println(table)
		}

		if format == "json" {
			json, _ := table.JSON(2)
			fmt.Println(json)
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().StringP("type", "t", "pod", "assign workload type (pod、deploy、sts、ds). default pod")
	imageCmd.Flags().StringP("namespace", "n", "default", "assign workload namespace")
	imageCmd.Flags().StringP("format", "f", "table", "show result fmt (table、json). default table")
	imageCmd.Flags().StringP("name", "c", "", "assign workload name")
	imageCmd.Flags().BoolP("all", "A", false, "show all namespace workload image")
}
