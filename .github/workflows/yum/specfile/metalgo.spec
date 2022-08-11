%define _build_id_links none

Name:           metalgo
Version:        %{version}
Release:        %{release}
Summary:        The Metal platform binaries
URL:            https://github.com/MetalBlockchain/%{name}
License:        BSD-3
AutoReqProv:    no

%description
Metal is an incredibly lightweight protocol, so the minimum computer requirements are quite modest.

%files
/usr/local/bin/metalgo
/usr/local/lib/metalgo
/usr/local/lib/metalgo/evm

%changelog
* Mon Oct 26 2020 Charlie Wyse <charlie@avalabs.org>
- First creation of package

