default['clang']               = {}
default['clang']['deb']        =
  'https://storage.googleapis.com/pl-infra-dev-artifacts/clang-9.0-pl5.deb'
default['clang']['deb_sha256'] =
  '6f261139154d4ffff9b3b5b2f049ab5e729df8d4fee1f55ffdb716a5d23a7e22'
default['clang']['version'] = "9.0-pl3"

default['gperftools']               = {}
default['gperftools']['deb']        =
  'https://storage.googleapis.com/pl-infra-dev-artifacts/gperftools-pixie-2.7-pl2.deb'
default['gperftools']['deb_sha256'] =
  'f43a343a6eae52dfd9ef1a3e3a9fe14347587d05fa7397f9444b473a4a65e959'
default['gperftools']['version'] = "2.7-pl2"


default['skaffold']                  = {}
default['kubectl']                   = {}
default['bazel']                     = {}
default['golang']                    = {}
default['sops']                      = {}
default['shellcheck']                = {}
default['sentry']                    = {}

if node[:platform] == 'ubuntu'
  default['bazel']['download_path'] =
    'https://github.com/bazelbuild/bazel/releases/download/2.2.0/bazel-2.2.0-linux-x86_64'
  default['bazel']['sha256'] =
    'b2f002ea0e6194a181af6ac84cd94bd8dc797722eb2354690bebac92dda233ff'

  default['golang']['download_path'] =
    'https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz'
  default['golang']['sha256'] =
    '0567734d558aef19112f2b2873caa0c600f1b4a5827930eb5a7f35235219e9d8'

  default['skaffold']['download_path'] =
    'https://storage.googleapis.com/pl-infra-dev-artifacts/skaffold/84eafe1c/skaffold-linux-amd64'
  default['skaffold']['sha256']        =
    '40b678384fb26c5b429160bf9d4b5cf60a619ce3f6b5d771335bde74df123024'

  default['kubectl']['download_path'] =
    'https://storage.googleapis.com/kubernetes-release/release/v1.14.6/bin/linux/amd64/kubectl'
  default['kubectl']['sha256']        = '5f8e8d8de929f64b8f779d0428854285e1a1c53a02cc2ad6b1ce5d32eefad25c'

  default['minikube']['download_path'] =
    'https://storage.googleapis.com/minikube/releases/v1.0.0/minikube-linux-amd64'
  default['minikube']['sha256']        =
    'a315869f81aae782ecc6ff2a6de4d0ab3a17ca1840d1d8e6eea050a8dd05907f'

  default['nodejs']['download_path'] = 'https://nodejs.org/dist/v12.16.1/node-v12.16.1-linux-x64.tar.gz'
  default['nodejs']['sha256']        = 'b2d9787da97d6c0d5cbf24c69fdbbf376b19089f921432c5a61aa323bc070bea'

  default['packer']['download_path'] = 'https://releases.hashicorp.com/packer/1.3.2/packer_1.3.2_linux_amd64.zip'
  default['packer']['sha256']        = '5e51808299135fee7a2e664b09f401b5712b5ef18bd4bad5bc50f4dcd8b149a1'

  default['kustomize']['download_path'] =
    'https://github.com/kubernetes-sigs/kustomize/releases/download/v3.2.0/kustomize_3.2.0_linux_amd64'
  default['kustomize']['sha256']        =
    '7db89e32575d81393d5d84f0dc6cbe444457e61ce71af06c6e6b7b6718299c22'

  default['sops']['download_path'] =
    'https://github.com/mozilla/sops/releases/download/3.3.1/sops-3.3.1.linux'
  default['sops']['sha256']        =
    '6eacdd01b68fd140eb71bbca233bea897cccb75dbf9e00a02e648b2f9a8a6939'

  default['shellcheck']['download_path'] =
    'https://storage.googleapis.com/shellcheck/shellcheck-v0.7.0.linux.x86_64.tar.xz'
  default['shellcheck']['sha256']        =
    '39c501aaca6aae3f3c7fc125b3c3af779ddbe4e67e4ebdc44c2ae5cba76c847f'

  default['sentry']['download_path'] =
    'https://github.com/getsentry/sentry-cli/releases/download/1.52.0/sentry-cli-Linux-x86_64'
  default['sentry']['sha256']        =
    'd6aeb45efbcdd3ec780f714b5082046ea1db31ff60ed0fc39916bbc8b6d708be'
elsif node[:platform] == 'mac_os_x'
  default['bazel']['download_path'] =
    'https://github.com/bazelbuild/bazel/releases/download/2.2.0/bazel-2.2.0-darwin-x86_64'
  default['bazel']['sha256'] =
    '6c786ab4ebc72afe1ba49341b3c8b2dbee7baa31a1d8b9a546b98b1307d220ab'

  default['golang']['download_path'] =
    'https://dl.google.com/go/go1.13.8.darwin-amd64.tar.gz'
  default['golang']['sha256'] =
    'e7bad54950e1d18c716ac9202b5406e7d4aca9aa4ca9e334a9742f75c2167a9c'

  default['skaffold']['download_path'] =
    'https://storage.googleapis.com/skaffold/releases/v1.3.1/skaffold-darwin-amd64'
  default['skaffold']['sha256']        = 'e2f9eb8277e95db95384e2093afbb5bd571347ea065689d8cfd54bae2b301b9a'

  default['kubectl']['download_path'] =
    'https://storage.googleapis.com/kubernetes-release/release/v1.14.6/bin/darwin/amd64/kubectl'
  default['kubectl']['sha256']        = 'de42dd22f67c135b749c75f389c70084c3fe840e3d89a03804edd255ac6ee829'

  default['minikube']['download_path'] =
    'https://storage.googleapis.com/minikube/releases/v1.0.0/minikube-darwin-amd64'
  default['minikube']['sha256']        = '865bd3a13c1ad3b7732b2bea35b26fef150f2b3cbfc257c5d1835527d1b331e9'

  default['nodejs']['download_path'] = 'https://nodejs.org/dist/v12.16.1/node-v12.16.1-darwin-x64.tar.gz'
  default['nodejs']['sha256']        = '34895bce210ca4b3cf19cd480e6563588880dd7f5d798f3782e3650580d35920'

  default['packer']['download_path'] = 'https://releases.hashicorp.com/packer/1.3.2/packer_1.3.2_darwin_amd64.zip'
  default['packer']['sha256']        = '1c2433239d801b017def8e66bbff4be3e7700b70248261b0abff2cd9c980bf5b'

  default['kustomize']['download_path'] =
    'https://github.com/kubernetes-sigs/kustomize/releases/download/v3.2.0/kustomize_3.2.0_darwin_amd64'
  default['kustomize']['sha256']        =
    'c7991a79470a52a95f1fac33f588b76f64e597ac64b54106e452f3a8f642c62e'

  default['sops']['download_path'] =
    'https://github.com/mozilla/sops/releases/download/3.3.1/sops-3.3.1.darwin'
  default['sops']['sha256']        =
    '09bb5920ae609bdf041b74843e2d8211a7059847b21729fadfbd3c3e33e67d26'

  default['shellcheck']['download_path'] =
    'https://storage.googleapis.com/shellcheck/shellcheck-v0.7.0.darwin.x86_64.tar.xz'
  default['shellcheck']['sha256']        =
    'c4edf1f04e53a35c39a7ef83598f2c50d36772e4cc942fb08a1114f9d48e5380'

  default['sentry']['download_path'] =
    'https://github.com/getsentry/sentry-cli/releases/download/1.52.0/sentry-cli-Darwin-x86_64'
  default['sentry']['sha256']        =
    '97c9bafbcf87bd7dea4f1069fe18f8e8265de6f7eab20c62ca9299e0fa8c2af6'
end
