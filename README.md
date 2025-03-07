# LeanBalancer - GitHub Instructions

## 📌 Setup
```sh
git clone https://github.com/siddhu949/Load-Balancer.git
cd Load-Balancer
git config --global user.name "Your GitHub Username"
git config --global user.email "your-email@example.com"
```

## 🔄 Workflow
```sh
git pull origin main            # Get latest changes
git checkout -b feature-branch  # Create a new branch
git add .
git commit -m "Commit message"
git push origin feature-branch  # Push changes
```

## 🔀 Merging
```sh
git checkout main
git pull origin main
git merge feature-branch
git push origin main
```

## 🗑️ Cleanup
```sh
git branch -d feature-branch
git push origin --delete feature-branch
```

## ⚠️ Other Useful Commands
```sh
git status                     # Check status
git stash && git stash pop     # Save & restore changes
git reset --hard HEAD          # Reset changes
```
git push origin main -u
---

📌 **Tips:**
- Pull before making changes.
- Use clear commit messages.
- Resolve conflicts carefully.

🚀 Happy coding!
To run the Apache bench mark:
 .\ab.exe -n 5000 -c 100 http://localhost:8080/
 To run the prometheus:
  .\prometheus.exe --config.file=prometheus.yml
 To run the main (loadbalancer):
  go run ./cmd/server/main.go
 To run the backend server:
  go run backendX.go
  To run the grafana:
  .\grafana-server.exe
  
