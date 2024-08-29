//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/cert-management/pkg/certman2/apis/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CertManagerConfiguration)(nil), (*config.CertManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CertManagerConfiguration_To_config_CertManagerConfiguration(a.(*CertManagerConfiguration), b.(*config.CertManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CertManagerConfiguration)(nil), (*CertManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CertManagerConfiguration_To_v1alpha1_CertManagerConfiguration(a.(*config.CertManagerConfiguration), b.(*CertManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClientConnection)(nil), (*config.ClientConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClientConnection_To_config_ClientConnection(a.(*ClientConnection), b.(*config.ClientConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ClientConnection)(nil), (*ClientConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClientConnection_To_v1alpha1_ClientConnection(a.(*config.ClientConnection), b.(*ClientConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneClientConnection)(nil), (*config.ControlPlaneClientConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneClientConnection_To_config_ControlPlaneClientConnection(a.(*ControlPlaneClientConnection), b.(*config.ControlPlaneClientConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControlPlaneClientConnection)(nil), (*ControlPlaneClientConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControlPlaneClientConnection_To_v1alpha1_ControlPlaneClientConnection(a.(*config.ControlPlaneClientConnection), b.(*ControlPlaneClientConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerConfiguration)(nil), (*config.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(a.(*ControllerConfiguration), b.(*config.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfiguration)(nil), (*ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*config.ControllerConfiguration), b.(*ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNSClientConnection)(nil), (*config.DNSClientConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSClientConnection_To_config_DNSClientConnection(a.(*DNSClientConnection), b.(*config.DNSClientConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DNSClientConnection)(nil), (*DNSClientConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DNSClientConnection_To_v1alpha1_DNSClientConnection(a.(*config.DNSClientConnection), b.(*DNSClientConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IssuerControllerConfig)(nil), (*config.IssuerControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IssuerControllerConfig_To_config_IssuerControllerConfig(a.(*IssuerControllerConfig), b.(*config.IssuerControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.IssuerControllerConfig)(nil), (*IssuerControllerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_IssuerControllerConfig_To_v1alpha1_IssuerControllerConfig(a.(*config.IssuerControllerConfig), b.(*IssuerControllerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Server)(nil), (*config.Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Server_To_config_Server(a.(*Server), b.(*config.Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Server)(nil), (*Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Server_To_v1alpha1_Server(a.(*config.Server), b.(*Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServerConfiguration)(nil), (*config.ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(a.(*ServerConfiguration), b.(*config.ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServerConfiguration)(nil), (*ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(a.(*config.ServerConfiguration), b.(*ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CertManagerConfiguration_To_config_CertManagerConfiguration(in *CertManagerConfiguration, out *config.CertManagerConfiguration, s conversion.Scope) error {
	if in.ClientConnection != nil {
		in, out := &in.ClientConnection, &out.ClientConnection
		*out = new(config.ClientConnection)
		if err := Convert_v1alpha1_ClientConnection_To_config_ClientConnection(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ClientConnection = nil
	}
	if in.ControlPlaneClientConnection != nil {
		in, out := &in.ControlPlaneClientConnection, &out.ControlPlaneClientConnection
		*out = new(config.ControlPlaneClientConnection)
		if err := Convert_v1alpha1_ControlPlaneClientConnection_To_config_ControlPlaneClientConnection(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ControlPlaneClientConnection = nil
	}
	if in.DNSClientConnection != nil {
		in, out := &in.DNSClientConnection, &out.DNSClientConnection
		*out = new(config.DNSClientConnection)
		if err := Convert_v1alpha1_DNSClientConnection_To_config_DNSClientConnection(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DNSClientConnection = nil
	}
	if err := configv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	out.LogFormat = in.LogFormat
	if err := Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(componentbaseconfig.DebuggingConfiguration)
		if err := configv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Debugging = nil
	}
	if err := Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(&in.Controllers, &out.Controllers, s); err != nil {
		return err
	}
	out.Class = in.Class
	return nil
}

// Convert_v1alpha1_CertManagerConfiguration_To_config_CertManagerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_CertManagerConfiguration_To_config_CertManagerConfiguration(in *CertManagerConfiguration, out *config.CertManagerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_CertManagerConfiguration_To_config_CertManagerConfiguration(in, out, s)
}

func autoConvert_config_CertManagerConfiguration_To_v1alpha1_CertManagerConfiguration(in *config.CertManagerConfiguration, out *CertManagerConfiguration, s conversion.Scope) error {
	if in.ClientConnection != nil {
		in, out := &in.ClientConnection, &out.ClientConnection
		*out = new(ClientConnection)
		if err := Convert_config_ClientConnection_To_v1alpha1_ClientConnection(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ClientConnection = nil
	}
	if in.ControlPlaneClientConnection != nil {
		in, out := &in.ControlPlaneClientConnection, &out.ControlPlaneClientConnection
		*out = new(ControlPlaneClientConnection)
		if err := Convert_config_ControlPlaneClientConnection_To_v1alpha1_ControlPlaneClientConnection(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ControlPlaneClientConnection = nil
	}
	if in.DNSClientConnection != nil {
		in, out := &in.DNSClientConnection, &out.DNSClientConnection
		*out = new(DNSClientConnection)
		if err := Convert_config_DNSClientConnection_To_v1alpha1_DNSClientConnection(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.DNSClientConnection = nil
	}
	if err := configv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	out.LogFormat = in.LogFormat
	if err := Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(configv1alpha1.DebuggingConfiguration)
		if err := configv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Debugging = nil
	}
	if err := Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(&in.Controllers, &out.Controllers, s); err != nil {
		return err
	}
	out.Class = in.Class
	return nil
}

// Convert_config_CertManagerConfiguration_To_v1alpha1_CertManagerConfiguration is an autogenerated conversion function.
func Convert_config_CertManagerConfiguration_To_v1alpha1_CertManagerConfiguration(in *config.CertManagerConfiguration, out *CertManagerConfiguration, s conversion.Scope) error {
	return autoConvert_config_CertManagerConfiguration_To_v1alpha1_CertManagerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ClientConnection_To_config_ClientConnection(in *ClientConnection, out *config.ClientConnection, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnectionConfiguration, &out.ClientConnectionConfiguration, s); err != nil {
		return err
	}
	out.CacheResyncPeriod = (*v1.Duration)(unsafe.Pointer(in.CacheResyncPeriod))
	return nil
}

// Convert_v1alpha1_ClientConnection_To_config_ClientConnection is an autogenerated conversion function.
func Convert_v1alpha1_ClientConnection_To_config_ClientConnection(in *ClientConnection, out *config.ClientConnection, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClientConnection_To_config_ClientConnection(in, out, s)
}

func autoConvert_config_ClientConnection_To_v1alpha1_ClientConnection(in *config.ClientConnection, out *ClientConnection, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnectionConfiguration, &out.ClientConnectionConfiguration, s); err != nil {
		return err
	}
	out.CacheResyncPeriod = (*v1.Duration)(unsafe.Pointer(in.CacheResyncPeriod))
	return nil
}

// Convert_config_ClientConnection_To_v1alpha1_ClientConnection is an autogenerated conversion function.
func Convert_config_ClientConnection_To_v1alpha1_ClientConnection(in *config.ClientConnection, out *ClientConnection, s conversion.Scope) error {
	return autoConvert_config_ClientConnection_To_v1alpha1_ClientConnection(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneClientConnection_To_config_ControlPlaneClientConnection(in *ControlPlaneClientConnection, out *config.ControlPlaneClientConnection, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnectionConfiguration, &out.ClientConnectionConfiguration, s); err != nil {
		return err
	}
	out.CacheResyncPeriod = (*v1.Duration)(unsafe.Pointer(in.CacheResyncPeriod))
	return nil
}

// Convert_v1alpha1_ControlPlaneClientConnection_To_config_ControlPlaneClientConnection is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneClientConnection_To_config_ControlPlaneClientConnection(in *ControlPlaneClientConnection, out *config.ControlPlaneClientConnection, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneClientConnection_To_config_ControlPlaneClientConnection(in, out, s)
}

func autoConvert_config_ControlPlaneClientConnection_To_v1alpha1_ControlPlaneClientConnection(in *config.ControlPlaneClientConnection, out *ControlPlaneClientConnection, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnectionConfiguration, &out.ClientConnectionConfiguration, s); err != nil {
		return err
	}
	out.CacheResyncPeriod = (*v1.Duration)(unsafe.Pointer(in.CacheResyncPeriod))
	return nil
}

// Convert_config_ControlPlaneClientConnection_To_v1alpha1_ControlPlaneClientConnection is an autogenerated conversion function.
func Convert_config_ControlPlaneClientConnection_To_v1alpha1_ControlPlaneClientConnection(in *config.ControlPlaneClientConnection, out *ControlPlaneClientConnection, s conversion.Scope) error {
	return autoConvert_config_ControlPlaneClientConnection_To_v1alpha1_ControlPlaneClientConnection(in, out, s)
}

func autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_IssuerControllerConfig_To_config_IssuerControllerConfig(&in.Issuer, &out.Issuer, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	if err := Convert_config_IssuerControllerConfig_To_v1alpha1_IssuerControllerConfig(&in.Issuer, &out.Issuer, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_DNSClientConnection_To_config_DNSClientConnection(in *DNSClientConnection, out *config.DNSClientConnection, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnectionConfiguration, &out.ClientConnectionConfiguration, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_DNSClientConnection_To_config_DNSClientConnection is an autogenerated conversion function.
func Convert_v1alpha1_DNSClientConnection_To_config_DNSClientConnection(in *DNSClientConnection, out *config.DNSClientConnection, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSClientConnection_To_config_DNSClientConnection(in, out, s)
}

func autoConvert_config_DNSClientConnection_To_v1alpha1_DNSClientConnection(in *config.DNSClientConnection, out *DNSClientConnection, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnectionConfiguration, &out.ClientConnectionConfiguration, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_DNSClientConnection_To_v1alpha1_DNSClientConnection is an autogenerated conversion function.
func Convert_config_DNSClientConnection_To_v1alpha1_DNSClientConnection(in *config.DNSClientConnection, out *DNSClientConnection, s conversion.Scope) error {
	return autoConvert_config_DNSClientConnection_To_v1alpha1_DNSClientConnection(in, out, s)
}

func autoConvert_v1alpha1_IssuerControllerConfig_To_config_IssuerControllerConfig(in *IssuerControllerConfig, out *config.IssuerControllerConfig, s conversion.Scope) error {
	out.ConcurrentSyncs = (*int)(unsafe.Pointer(in.ConcurrentSyncs))
	out.SyncPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncPeriod))
	out.Namespace = in.Namespace
	out.DefaultIssuerName = in.DefaultIssuerName
	out.DefaultRequestsPerDayQuota = in.DefaultRequestsPerDayQuota
	return nil
}

// Convert_v1alpha1_IssuerControllerConfig_To_config_IssuerControllerConfig is an autogenerated conversion function.
func Convert_v1alpha1_IssuerControllerConfig_To_config_IssuerControllerConfig(in *IssuerControllerConfig, out *config.IssuerControllerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_IssuerControllerConfig_To_config_IssuerControllerConfig(in, out, s)
}

func autoConvert_config_IssuerControllerConfig_To_v1alpha1_IssuerControllerConfig(in *config.IssuerControllerConfig, out *IssuerControllerConfig, s conversion.Scope) error {
	out.ConcurrentSyncs = (*int)(unsafe.Pointer(in.ConcurrentSyncs))
	out.SyncPeriod = (*v1.Duration)(unsafe.Pointer(in.SyncPeriod))
	out.Namespace = in.Namespace
	out.DefaultIssuerName = in.DefaultIssuerName
	out.DefaultRequestsPerDayQuota = in.DefaultRequestsPerDayQuota
	return nil
}

// Convert_config_IssuerControllerConfig_To_v1alpha1_IssuerControllerConfig is an autogenerated conversion function.
func Convert_config_IssuerControllerConfig_To_v1alpha1_IssuerControllerConfig(in *config.IssuerControllerConfig, out *IssuerControllerConfig, s conversion.Scope) error {
	return autoConvert_config_IssuerControllerConfig_To_v1alpha1_IssuerControllerConfig(in, out, s)
}

func autoConvert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_v1alpha1_Server_To_config_Server is an autogenerated conversion function.
func Convert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	return autoConvert_v1alpha1_Server_To_config_Server(in, out, s)
}

func autoConvert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_config_Server_To_v1alpha1_Server is an autogenerated conversion function.
func Convert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	return autoConvert_config_Server_To_v1alpha1_Server(in, out, s)
}

func autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_Server_To_config_Server(&in.Webhooks, &out.Webhooks, s); err != nil {
		return err
	}
	out.HealthProbes = (*config.Server)(unsafe.Pointer(in.HealthProbes))
	out.Metrics = (*config.Server)(unsafe.Pointer(in.Metrics))
	return nil
}

// Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in, out, s)
}

func autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	if err := Convert_config_Server_To_v1alpha1_Server(&in.Webhooks, &out.Webhooks, s); err != nil {
		return err
	}
	out.HealthProbes = (*Server)(unsafe.Pointer(in.HealthProbes))
	out.Metrics = (*Server)(unsafe.Pointer(in.Metrics))
	return nil
}

// Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration is an autogenerated conversion function.
func Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in, out, s)
}
