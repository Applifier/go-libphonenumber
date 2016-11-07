#! /usr/bin/env bash

# See https://github.com/googlei18n/libphonenumber/tree/master/cpp

# APT deps
sudo apt-get update -qq
sudo apt-get install -y \
      cmake cmake-curses-gui libprotobuf-dev \
      libicu-dev libboost-dev libboost-thread-dev libboost-system-dev \
      protobuf-compiler sudo libc-bin libc-dev-bin libc6-dev clang cmake llvm

# libre2 cannot be found in ubuntu apt
cd ~
wget https://github.com/google/re2/archive/2016-04-01.tar.gz
tar -xzf 2016-04-01.tar.gz
cd re2-2016-04-01
make
sudo make install

# travis libgtest is too old
cd ~
git clone https://github.com/google/googletest.git

# libphonenumber itself
cd ~
wget https://github.com/googlei18n/libphonenumber/archive/libphonenumber-${LIBPHONENUMBER_VERSION}.tar.gz
tar -xzf libphonenumber-${LIBPHONENUMBER_VERSION}.tar.gz
cd libphonenumber-libphonenumber-${LIBPHONENUMBER_VERSION}/cpp
mkdir build
cd build
cmake \
  -DGTEST_SOURCE_DIR=~/googletest/googletest/ \
  -DGTEST_INCLUDE_DIR=~/googletest/googletest/include/ \
  ..
make
sudo make install
sudo ldconfig
