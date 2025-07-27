# `/build`

打包和持续集成。

将云（AMI），容器（Docker），OS（DEB，RPM，PKG）放置在 `/build/package` 目录中。

将您的CI（Travis，Circle，无人机）配置和脚本放在 `/build/ci` 目录中。
请注意，某些CI工具（例如Travis CI）对其配置文件的位置非常挑剔。
尝试将配置文件放入 `/build/ci` 目录中，将它们链接到CI工具在可能的情况下期望它们的位置（如果可能的话）。
