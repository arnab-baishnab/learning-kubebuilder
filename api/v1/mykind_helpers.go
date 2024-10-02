package v1

import "strings"

func (b *MyKind) DeploymentName() string {
	if b.Spec.DeploymentName != "" {
		return b.Spec.DeploymentName
	}
	return strings.Join([]string{b.Name, "dep"}, "-")
}

func (b *MyKind) ServiceName() string { return strings.Join([]string{b.Name, "svc"}, "-") }
