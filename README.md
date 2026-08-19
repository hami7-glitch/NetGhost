# NetGhost 🛡️

NetGhost is an experimental network monitoring and anomaly detection project.

## 🎯 Goal

The goal of NetGhost is to monitor network activity and identify potentially suspicious behavior such as:

- Port scanning
- Unusual connection patterns
- Network anomalies

## 🚧 Project Status

NetGhost is currently in early development.

Version: `v0.1`

The first version will focus on a simple and understandable network monitoring architecture before adding advanced technologies such as eBPF, machine learning, databases, or a web dashboard.

## 🧠 Planned Architecture

```text
Network Activity
       │
       ▼
   Collector
       │
       ▼
   Network Events
       │
       ▼
 Detection Rules
       │
   ┌───┴───┐
   ▼       ▼
Normal   Suspicious
           │
           ▼
         Alert

🛠️ Planned Technologies
Go
Linux
eBPF
Network monitoring
Rule-based anomaly detection
Advanced components such as machine learning, databases, and web dashboards may be added in later versions.
📌 Development
NetGhost is being developed incrementally, starting from a minimal working prototype.
The project is intended to be educational and experimental.
