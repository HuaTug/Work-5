CURDIR=$(cd $(dirname $0); pwd)
# $()：这是一个命令替换操作符。它会执行括号内的命令，并将结果替换到原来的位置。因此，$(dirname $0) 会被替换为当前脚本所在的目录。
# $0：这是一个特殊的 shell 变量，代表当前脚本的名称。如果你在命令行中直接运行一个脚本，那么 $0 就是该脚本的路径
# 第一个$表示为命令替换，即将$所包含的命令使用一个变量替代 
BinaryName=test
echo $CURDIR/bin/$BinaryName
exec $CURDIR/bin/$BinaryName