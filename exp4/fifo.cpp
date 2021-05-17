#include "header.h"

FIFO::FIFO(int capacity)
{
    this->capacity = capacity;
    this->count = 0;
    this->missCount = 0;
}

void FIFO::get(int k)
{
    this->count++;
    // 没有命中
    if (!cache.count(k)) {
        this->missCount++;
        // 容量满，先删除最早加入的
        if (cache.size() == this->capacity) {
            auto key = queue.front();
            cache.erase(key);
            queue.pop();
        }
        queue.push(k);
        cache[k] = true;
    }
}