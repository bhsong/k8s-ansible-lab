# 🚀 Raspberry Pi Kubernetes Lab (K3s + Ansible + Go)

**라즈베리파이(ARM64)** 환경에서 **Kubernetes(K3s)** 클러스터를 구축하고, **Go 웹 애플리케이션**을 배포하는 풀 스택 DevOps 프로젝트입니다.

모든 인프라 구축 과정은 **Ansible**로 자동화되어 있으며, **GitHub Actions**를 통해 CI/CD 파이프라인이 구축되어 있습니다.

## 🛠 Tech Stack

- **Infrastructure:** Raspberry Pi 4/5, K3s (Lightweight Kubernetes)
- **Automation:** Ansible (IaC)
- **CI/CD:** GitHub Actions, GHCR (Container Registry)
- **App:** Go (Golang), PostgreSQL
- **Network:** MetalLB (or K3s default), Ingress (Traefik), Self-signed SSL
- **Container:** Docker (Multi-arch build for ARM64)

## 📂 Directory Structure

```bash
.
├── .github/workflows  # GitHub Actions CI/CD Script
├── ansible            # Ansible Playbooks for Infrastructure setup
├── app                # Go Web Application Source Code
├── k8s                # Kubernetes Manifests (Deployment, Service, Ingress, PVC)
└── README.md
```

## ✨ Features
1. Automated K3s Setup: Ansible을 이용한 원클릭 K3s 설치 및 설정.

2. Multi-Arch Build: GitHub Actions QEMU를 이용해 ARM64/AMD64 이미지 동시 빌드.

3. Database Integration: PostgreSQL StatefulSet 구성 및 PVC를 통한 데이터 영속성 보장.

4. Secure Access: Ingress 및 Self-signed Certificate를 이용한 HTTPS(SSL) 적용.

5. Traffic Management: 방문자 수 카운팅 로직 구현.

## 🚀 How to Run
1. Prerequisites
- Raspberry Pi (Ubuntu or Raspberry Pi OS 64bit)
- Git installed

2. Clone Repository
```Bash
git clone [https://github.com/bhsong/k8s-ansible-lab.git](https://github.com/bhsong/k8s-ansible-lab.git)
cd k8s-ansible-lab
```

3. Configure Inventory
Create an inventory file from the example.
```Bash
cp ansible/inventory.example.ini ansible/inventory.ini
# Edit inventory.ini with your IP
```

4. Run Ansible Playbook
This command will install K3s, generate SSL certs, and deploy the full stack.
```Bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

5. Verify
Add the domain to your local hosts file (/etc/hosts or C:\Windows\System32\drivers\etc\hosts).
```Plaintext
[Raspberry_Pi_IP]  k8s-lab.com
```

Access via browser:
- https://www.google.com/search?q=https://k8s-lab.com

## 📝 License
This project is licensed under the MIT License.