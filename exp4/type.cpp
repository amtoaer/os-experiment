#include <header.h>

DoubleList::DoubleList()
{
    this->begin = &Node();
    this->end = &Node();
    this->begin->next = this->end;
    this->end->next = this->begin;
}

Node* DoubleList::insertNode(int key)
{
    auto tmp = this->begin->next;
    auto newNode = &Node(key, this->begin, tmp);
    this->begin->next = newNode;
    tmp->prev = newNode;
    return newNode;
}

void DoubleList::deleteNode(Node* node)
{
    auto prev = node->prev;
    auto next = node->next;
    prev->next = next;
    next->prev = prev;
}

Node* DoubleList::getLast()
{
    return this->end->prev;
}

void DoubleList::moveToFirst(Node* node)
{
    this->deleteNode(node);
    auto tmp = this->begin->next;
    node->prev = this->begin;
    node->next = tmp;
    tmp->prev = node;
    this->begin->next = tmp;
}

LRU::LRU(int capacity)
{
    this->d = DoubleList();
    this->capacity = capacity;
}

void LRU::get(int k)
{
    this->count++;
    // 如果已经存在
    if (cache.count(k)) {
        // 将该元素调到双向链表头部
        auto node = cache.at(k);
        this->d.moveToFirst(node);
    } else {
        // 发生缺页
        this->missCount++;
        // 如果已经满了，先删除掉最久没有使用的
        if (cache.size() == this->capacity) {
            auto last = this->d.getLast();
            this->d.deleteNode(last);
            cache.erase(last->key);
        }
        auto node = this->d.insertNode(k);
        this->cache[k] = node;
    }
}

float LRU::getMissingRate()
{
    return float(this->missCount) / this->count;
}