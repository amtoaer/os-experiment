#include "header.h"
#include <iostream>

int main()
{
    auto fifo = FIFO(3);
    auto lru = LRU(3);
    int testCase[] = { 1, 2, 3, 4, 2, 1, 5, 6, 2, 1, 2, 3, 7, 6, 3, 2, 1, 2, 3, 6 };
    for (int i : testCase) {
        fifo.get(i);
        lru.get(i);
    }
    std::cout << fifo.getCount() << std::endl;
    std::cout << fifo.getMissCount() << std::endl;
    std::cout << fifo.getMissingRate() << std::endl;
    std::cout << lru.getCount() << std::endl;
    std::cout << lru.getMissCount() << std::endl;
    std::cout << lru.getMissingRate() << std::endl;
}