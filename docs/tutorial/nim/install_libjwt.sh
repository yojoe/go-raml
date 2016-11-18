cd /tmp
sudo apt-get install libtool autoconf libjansson-dev
git clone https://github.com/benmcollins/libjwt.git
cd libjwt
autoreconf -i
./configure
make
sudo make install
sudo ldconfig
