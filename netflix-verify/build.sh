#!/bin/bash

machine_arch=$(uname -m)

if [ "$machine_arch" == "x86_64" ]; then
    echo "处理器架构为 x86"
elif [ "$machine_arch" == "arm64" ]; then
    echo "处理器架构为 M1"
else
    echo "无法确定处理器架构"
fi

if [ ! -z "$machine_arch" ]; then
    odir="./output/$machine_arch/nf"
    echo "$odir"
    go build -o "$odir"
    echo "編譯完畢: $odir"
fi
