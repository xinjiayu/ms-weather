FROM alpine # 把项目运行在最小化linux系统上
ADD example-srv /example-srv # 把已经build好的example-srv可执行文件拷贝到docker 容器内相应目录下
ENTRYPOINT [ "/example-srv" ]   # 直接运行 这里就是运行shell命令
CMD [ "$0" "$@"]	# cmd 在没有ENTRYPOINT 的情况下可以做为执行命令的入口,但现在已经有entrypoint的情况下当做传参,
