# 极简版抖音后端

- 接口说明：[视频流接口 - 极简版抖音 (apifox.cn)](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
- 实现接口：
  - 基础接口
  - 互动接口
- 技术栈：
  - hertz
  - kitex
  - mysql+gorm
  - kafka
  - redis

- 构建和运行：

  ```shell
  docker-compose up
  
  make
  
  cd cmd/api
  go run .
  
  # run in a new terminal
  cd cmd/feed
  go run .
  
  # run in a new terminal
  cd cmd/user
  go run .
  
  # run in a new terminal
  cd cmd/publish
  go run .
  
  # run in a new terminal
  cd cmd/favorite
  go run .
  
  # run in a new terminal
  cd cmd/comment
  go run .
  ```

  