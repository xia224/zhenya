# Keep your linux enviroment running as smoothly and effectively as possible.

## Tip1 禁用不必要的组件
Linux系统启动时，可选择启动的组件和后台服务，那些不需要的组件会耗费额外的RAM和CPU资源。
最好的办法是在系统启动脚本里禁用它们。

## Tip2 系统更新
每个新的稳定发行版都会带有新的bug fix和安全补丁，所以建议常更新系统到最新的稳定版本。
当然了，建议先在沙箱里测验新系统是否存在潜在的问题，然后再更新生产环境的系统。

## Tip3 调整TCP linux size
提高数据传输率。
如何调整？ 待补充
CPU 使用率、虚拟内存、kernel 状态统计、网络I/O、网络错误统计、磁盘I/O


## Tip4 优化虚拟机、容器
目前，开发者会在自己的windows或macos上启动虚拟机或者docker容器，隔离性比较好，类似沙箱，也可以启动不同的
操作系统环境。
其实，禁用不必要的服务、优化性能、减轻负载以及屏蔽广告等等，有很多措施优化你的虚拟机或容器。

## Tip5 为mysql等数据库和apache开源软件设定合理的配置
linux基础系统不是孤立的，它和其上运行的各种服务有关联，例如mysql和apache为了获取更高的性能，应该也被优化过，如何mysql可以通过调整cache大小，获取更多的RAM。
所以，合理的配置设置，能帮助linux系统节省RAM。

## Tip6 熟知5个常用的linux性能查看命令
top：查看进程/线程使用情况
vmstat: 查看虚拟内存统计
iostat: 查看I/O统计
free: 查看内存使用情况
sar: 找出系统瓶颈的利器
更多的命令，请参看[linux工具](https://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/sar.html)

