# 蓝绿发布


## blue green

```bash
  strategy:
    blueGreen:
      # 活跃的 Service
      activeService: shark

      # 预灰度的 Service
      previewService: shark-canary

      # 蓝绿切换
      autoPromotionEnabled: true

```
