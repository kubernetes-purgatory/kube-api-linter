package c

type OptionalOrRequiredTestStruct struct {
	RequiredEnumField RequiredEnum // want "field RequiredEnumField must be marked as kubebuilder:validation:Optional or kubebuilder:validation:Required"

	KubebuilderRequiredEnumField KubeBuilderRequiredEnum // want "field KubebuilderRequiredEnumField must be marked as kubebuilder:validation:Optional or kubebuilder:validation:Required"

	OptionalEnumField OptionalEnum // want "field OptionalEnumField must be marked as kubebuilder:validation:Optional or kubebuilder:validation:Required"

	KubebuilderOptionalEnumField KubeBuilderOptionalEnum // want "field KubebuilderOptionalEnumField must be marked as kubebuilder:validation:Optional or kubebuilder:validation:Required"
}

// +required
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type RequiredEnum string // want "required should not be defined on a type RequiredEnum"

// +kubebuilder:validation:Required
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type KubeBuilderRequiredEnum string // want "kubebuilder:validation:Required should not be defined on a type KubeBuilderRequiredEnum"

// +optional
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type OptionalEnum string // want "optional should not be defined on a type OptionalEnum"

// +kubebuilder:validation:Optional
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type KubeBuilderOptionalEnum string // want "kubebuilder:validation:Optional should not be defined on a type KubeBuilderOptionalEnum"
