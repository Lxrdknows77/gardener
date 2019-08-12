// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	settings "github.com/gardener/gardener/pkg/apis/settings"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterOpenIDConnectPreset)(nil), (*settings.ClusterOpenIDConnectPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClusterOpenIDConnectPreset_To_settings_ClusterOpenIDConnectPreset(a.(*ClusterOpenIDConnectPreset), b.(*settings.ClusterOpenIDConnectPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.ClusterOpenIDConnectPreset)(nil), (*ClusterOpenIDConnectPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_ClusterOpenIDConnectPreset_To_v1alpha1_ClusterOpenIDConnectPreset(a.(*settings.ClusterOpenIDConnectPreset), b.(*ClusterOpenIDConnectPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterOpenIDConnectPresetList)(nil), (*settings.ClusterOpenIDConnectPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClusterOpenIDConnectPresetList_To_settings_ClusterOpenIDConnectPresetList(a.(*ClusterOpenIDConnectPresetList), b.(*settings.ClusterOpenIDConnectPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.ClusterOpenIDConnectPresetList)(nil), (*ClusterOpenIDConnectPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_ClusterOpenIDConnectPresetList_To_v1alpha1_ClusterOpenIDConnectPresetList(a.(*settings.ClusterOpenIDConnectPresetList), b.(*ClusterOpenIDConnectPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeAPIServerOpenIDConnect)(nil), (*settings.KubeAPIServerOpenIDConnect)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeAPIServerOpenIDConnect_To_settings_KubeAPIServerOpenIDConnect(a.(*KubeAPIServerOpenIDConnect), b.(*settings.KubeAPIServerOpenIDConnect), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.KubeAPIServerOpenIDConnect)(nil), (*KubeAPIServerOpenIDConnect)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_KubeAPIServerOpenIDConnect_To_v1alpha1_KubeAPIServerOpenIDConnect(a.(*settings.KubeAPIServerOpenIDConnect), b.(*KubeAPIServerOpenIDConnect), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OpenIDConnectClientAuthentication)(nil), (*settings.OpenIDConnectClientAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OpenIDConnectClientAuthentication_To_settings_OpenIDConnectClientAuthentication(a.(*OpenIDConnectClientAuthentication), b.(*settings.OpenIDConnectClientAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.OpenIDConnectClientAuthentication)(nil), (*OpenIDConnectClientAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_OpenIDConnectClientAuthentication_To_v1alpha1_OpenIDConnectClientAuthentication(a.(*settings.OpenIDConnectClientAuthentication), b.(*OpenIDConnectClientAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OpenIDConnectPreset)(nil), (*settings.OpenIDConnectPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OpenIDConnectPreset_To_settings_OpenIDConnectPreset(a.(*OpenIDConnectPreset), b.(*settings.OpenIDConnectPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.OpenIDConnectPreset)(nil), (*OpenIDConnectPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_OpenIDConnectPreset_To_v1alpha1_OpenIDConnectPreset(a.(*settings.OpenIDConnectPreset), b.(*OpenIDConnectPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OpenIDConnectPresetList)(nil), (*settings.OpenIDConnectPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OpenIDConnectPresetList_To_settings_OpenIDConnectPresetList(a.(*OpenIDConnectPresetList), b.(*settings.OpenIDConnectPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.OpenIDConnectPresetList)(nil), (*OpenIDConnectPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_OpenIDConnectPresetList_To_v1alpha1_OpenIDConnectPresetList(a.(*settings.OpenIDConnectPresetList), b.(*OpenIDConnectPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OpenIDConnectPresetSpec)(nil), (*settings.OpenIDConnectPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec(a.(*OpenIDConnectPresetSpec), b.(*settings.OpenIDConnectPresetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.OpenIDConnectPresetSpec)(nil), (*OpenIDConnectPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec(a.(*settings.OpenIDConnectPresetSpec), b.(*OpenIDConnectPresetSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ClusterOpenIDConnectPreset_To_settings_ClusterOpenIDConnectPreset(in *ClusterOpenIDConnectPreset, out *settings.ClusterOpenIDConnectPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec(&in.OpenIDConnectPresetSpec, &out.OpenIDConnectPresetSpec, s); err != nil {
		return err
	}
	out.ProjectSelector = (*v1.LabelSelector)(unsafe.Pointer(in.ProjectSelector))
	return nil
}

// Convert_v1alpha1_ClusterOpenIDConnectPreset_To_settings_ClusterOpenIDConnectPreset is an autogenerated conversion function.
func Convert_v1alpha1_ClusterOpenIDConnectPreset_To_settings_ClusterOpenIDConnectPreset(in *ClusterOpenIDConnectPreset, out *settings.ClusterOpenIDConnectPreset, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterOpenIDConnectPreset_To_settings_ClusterOpenIDConnectPreset(in, out, s)
}

func autoConvert_settings_ClusterOpenIDConnectPreset_To_v1alpha1_ClusterOpenIDConnectPreset(in *settings.ClusterOpenIDConnectPreset, out *ClusterOpenIDConnectPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec(&in.OpenIDConnectPresetSpec, &out.OpenIDConnectPresetSpec, s); err != nil {
		return err
	}
	out.ProjectSelector = (*v1.LabelSelector)(unsafe.Pointer(in.ProjectSelector))
	return nil
}

// Convert_settings_ClusterOpenIDConnectPreset_To_v1alpha1_ClusterOpenIDConnectPreset is an autogenerated conversion function.
func Convert_settings_ClusterOpenIDConnectPreset_To_v1alpha1_ClusterOpenIDConnectPreset(in *settings.ClusterOpenIDConnectPreset, out *ClusterOpenIDConnectPreset, s conversion.Scope) error {
	return autoConvert_settings_ClusterOpenIDConnectPreset_To_v1alpha1_ClusterOpenIDConnectPreset(in, out, s)
}

func autoConvert_v1alpha1_ClusterOpenIDConnectPresetList_To_settings_ClusterOpenIDConnectPresetList(in *ClusterOpenIDConnectPresetList, out *settings.ClusterOpenIDConnectPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]settings.ClusterOpenIDConnectPreset)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ClusterOpenIDConnectPresetList_To_settings_ClusterOpenIDConnectPresetList is an autogenerated conversion function.
func Convert_v1alpha1_ClusterOpenIDConnectPresetList_To_settings_ClusterOpenIDConnectPresetList(in *ClusterOpenIDConnectPresetList, out *settings.ClusterOpenIDConnectPresetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterOpenIDConnectPresetList_To_settings_ClusterOpenIDConnectPresetList(in, out, s)
}

func autoConvert_settings_ClusterOpenIDConnectPresetList_To_v1alpha1_ClusterOpenIDConnectPresetList(in *settings.ClusterOpenIDConnectPresetList, out *ClusterOpenIDConnectPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterOpenIDConnectPreset)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_settings_ClusterOpenIDConnectPresetList_To_v1alpha1_ClusterOpenIDConnectPresetList is an autogenerated conversion function.
func Convert_settings_ClusterOpenIDConnectPresetList_To_v1alpha1_ClusterOpenIDConnectPresetList(in *settings.ClusterOpenIDConnectPresetList, out *ClusterOpenIDConnectPresetList, s conversion.Scope) error {
	return autoConvert_settings_ClusterOpenIDConnectPresetList_To_v1alpha1_ClusterOpenIDConnectPresetList(in, out, s)
}

func autoConvert_v1alpha1_KubeAPIServerOpenIDConnect_To_settings_KubeAPIServerOpenIDConnect(in *KubeAPIServerOpenIDConnect, out *settings.KubeAPIServerOpenIDConnect, s conversion.Scope) error {
	out.CABundle = (*string)(unsafe.Pointer(in.CABundle))
	out.ClientID = in.ClientID
	out.GroupsClaim = (*string)(unsafe.Pointer(in.GroupsClaim))
	out.GroupsPrefix = (*string)(unsafe.Pointer(in.GroupsPrefix))
	out.IssuerURL = in.IssuerURL
	out.RequiredClaims = *(*map[string]string)(unsafe.Pointer(&in.RequiredClaims))
	out.SigningAlgs = *(*[]string)(unsafe.Pointer(&in.SigningAlgs))
	out.UsernameClaim = (*string)(unsafe.Pointer(in.UsernameClaim))
	out.UsernamePrefix = (*string)(unsafe.Pointer(in.UsernamePrefix))
	return nil
}

// Convert_v1alpha1_KubeAPIServerOpenIDConnect_To_settings_KubeAPIServerOpenIDConnect is an autogenerated conversion function.
func Convert_v1alpha1_KubeAPIServerOpenIDConnect_To_settings_KubeAPIServerOpenIDConnect(in *KubeAPIServerOpenIDConnect, out *settings.KubeAPIServerOpenIDConnect, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeAPIServerOpenIDConnect_To_settings_KubeAPIServerOpenIDConnect(in, out, s)
}

func autoConvert_settings_KubeAPIServerOpenIDConnect_To_v1alpha1_KubeAPIServerOpenIDConnect(in *settings.KubeAPIServerOpenIDConnect, out *KubeAPIServerOpenIDConnect, s conversion.Scope) error {
	out.CABundle = (*string)(unsafe.Pointer(in.CABundle))
	out.ClientID = in.ClientID
	out.GroupsClaim = (*string)(unsafe.Pointer(in.GroupsClaim))
	out.GroupsPrefix = (*string)(unsafe.Pointer(in.GroupsPrefix))
	out.IssuerURL = in.IssuerURL
	out.RequiredClaims = *(*map[string]string)(unsafe.Pointer(&in.RequiredClaims))
	out.SigningAlgs = *(*[]string)(unsafe.Pointer(&in.SigningAlgs))
	out.UsernameClaim = (*string)(unsafe.Pointer(in.UsernameClaim))
	out.UsernamePrefix = (*string)(unsafe.Pointer(in.UsernamePrefix))
	return nil
}

// Convert_settings_KubeAPIServerOpenIDConnect_To_v1alpha1_KubeAPIServerOpenIDConnect is an autogenerated conversion function.
func Convert_settings_KubeAPIServerOpenIDConnect_To_v1alpha1_KubeAPIServerOpenIDConnect(in *settings.KubeAPIServerOpenIDConnect, out *KubeAPIServerOpenIDConnect, s conversion.Scope) error {
	return autoConvert_settings_KubeAPIServerOpenIDConnect_To_v1alpha1_KubeAPIServerOpenIDConnect(in, out, s)
}

func autoConvert_v1alpha1_OpenIDConnectClientAuthentication_To_settings_OpenIDConnectClientAuthentication(in *OpenIDConnectClientAuthentication, out *settings.OpenIDConnectClientAuthentication, s conversion.Scope) error {
	out.Secret = (*string)(unsafe.Pointer(in.Secret))
	out.ExtraConfig = *(*map[string]string)(unsafe.Pointer(&in.ExtraConfig))
	return nil
}

// Convert_v1alpha1_OpenIDConnectClientAuthentication_To_settings_OpenIDConnectClientAuthentication is an autogenerated conversion function.
func Convert_v1alpha1_OpenIDConnectClientAuthentication_To_settings_OpenIDConnectClientAuthentication(in *OpenIDConnectClientAuthentication, out *settings.OpenIDConnectClientAuthentication, s conversion.Scope) error {
	return autoConvert_v1alpha1_OpenIDConnectClientAuthentication_To_settings_OpenIDConnectClientAuthentication(in, out, s)
}

func autoConvert_settings_OpenIDConnectClientAuthentication_To_v1alpha1_OpenIDConnectClientAuthentication(in *settings.OpenIDConnectClientAuthentication, out *OpenIDConnectClientAuthentication, s conversion.Scope) error {
	out.Secret = (*string)(unsafe.Pointer(in.Secret))
	out.ExtraConfig = *(*map[string]string)(unsafe.Pointer(&in.ExtraConfig))
	return nil
}

// Convert_settings_OpenIDConnectClientAuthentication_To_v1alpha1_OpenIDConnectClientAuthentication is an autogenerated conversion function.
func Convert_settings_OpenIDConnectClientAuthentication_To_v1alpha1_OpenIDConnectClientAuthentication(in *settings.OpenIDConnectClientAuthentication, out *OpenIDConnectClientAuthentication, s conversion.Scope) error {
	return autoConvert_settings_OpenIDConnectClientAuthentication_To_v1alpha1_OpenIDConnectClientAuthentication(in, out, s)
}

func autoConvert_v1alpha1_OpenIDConnectPreset_To_settings_OpenIDConnectPreset(in *OpenIDConnectPreset, out *settings.OpenIDConnectPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec(&in.OpenIDConnectPresetSpec, &out.OpenIDConnectPresetSpec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_OpenIDConnectPreset_To_settings_OpenIDConnectPreset is an autogenerated conversion function.
func Convert_v1alpha1_OpenIDConnectPreset_To_settings_OpenIDConnectPreset(in *OpenIDConnectPreset, out *settings.OpenIDConnectPreset, s conversion.Scope) error {
	return autoConvert_v1alpha1_OpenIDConnectPreset_To_settings_OpenIDConnectPreset(in, out, s)
}

func autoConvert_settings_OpenIDConnectPreset_To_v1alpha1_OpenIDConnectPreset(in *settings.OpenIDConnectPreset, out *OpenIDConnectPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec(&in.OpenIDConnectPresetSpec, &out.OpenIDConnectPresetSpec, s); err != nil {
		return err
	}
	return nil
}

// Convert_settings_OpenIDConnectPreset_To_v1alpha1_OpenIDConnectPreset is an autogenerated conversion function.
func Convert_settings_OpenIDConnectPreset_To_v1alpha1_OpenIDConnectPreset(in *settings.OpenIDConnectPreset, out *OpenIDConnectPreset, s conversion.Scope) error {
	return autoConvert_settings_OpenIDConnectPreset_To_v1alpha1_OpenIDConnectPreset(in, out, s)
}

func autoConvert_v1alpha1_OpenIDConnectPresetList_To_settings_OpenIDConnectPresetList(in *OpenIDConnectPresetList, out *settings.OpenIDConnectPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]settings.OpenIDConnectPreset)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_OpenIDConnectPresetList_To_settings_OpenIDConnectPresetList is an autogenerated conversion function.
func Convert_v1alpha1_OpenIDConnectPresetList_To_settings_OpenIDConnectPresetList(in *OpenIDConnectPresetList, out *settings.OpenIDConnectPresetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_OpenIDConnectPresetList_To_settings_OpenIDConnectPresetList(in, out, s)
}

func autoConvert_settings_OpenIDConnectPresetList_To_v1alpha1_OpenIDConnectPresetList(in *settings.OpenIDConnectPresetList, out *OpenIDConnectPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]OpenIDConnectPreset)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_settings_OpenIDConnectPresetList_To_v1alpha1_OpenIDConnectPresetList is an autogenerated conversion function.
func Convert_settings_OpenIDConnectPresetList_To_v1alpha1_OpenIDConnectPresetList(in *settings.OpenIDConnectPresetList, out *OpenIDConnectPresetList, s conversion.Scope) error {
	return autoConvert_settings_OpenIDConnectPresetList_To_v1alpha1_OpenIDConnectPresetList(in, out, s)
}

func autoConvert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec(in *OpenIDConnectPresetSpec, out *settings.OpenIDConnectPresetSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_KubeAPIServerOpenIDConnect_To_settings_KubeAPIServerOpenIDConnect(&in.Server, &out.Server, s); err != nil {
		return err
	}
	out.Client = (*settings.OpenIDConnectClientAuthentication)(unsafe.Pointer(in.Client))
	out.ShootSelector = (*v1.LabelSelector)(unsafe.Pointer(in.ShootSelector))
	out.Weight = in.Weight
	return nil
}

// Convert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec is an autogenerated conversion function.
func Convert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec(in *OpenIDConnectPresetSpec, out *settings.OpenIDConnectPresetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_OpenIDConnectPresetSpec_To_settings_OpenIDConnectPresetSpec(in, out, s)
}

func autoConvert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec(in *settings.OpenIDConnectPresetSpec, out *OpenIDConnectPresetSpec, s conversion.Scope) error {
	if err := Convert_settings_KubeAPIServerOpenIDConnect_To_v1alpha1_KubeAPIServerOpenIDConnect(&in.Server, &out.Server, s); err != nil {
		return err
	}
	out.Client = (*OpenIDConnectClientAuthentication)(unsafe.Pointer(in.Client))
	out.ShootSelector = (*v1.LabelSelector)(unsafe.Pointer(in.ShootSelector))
	out.Weight = in.Weight
	return nil
}

// Convert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec is an autogenerated conversion function.
func Convert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec(in *settings.OpenIDConnectPresetSpec, out *OpenIDConnectPresetSpec, s conversion.Scope) error {
	return autoConvert_settings_OpenIDConnectPresetSpec_To_v1alpha1_OpenIDConnectPresetSpec(in, out, s)
}
