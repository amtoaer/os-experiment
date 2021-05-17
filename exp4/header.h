#include <unordered_map>

// 单个节点
typedef struct Node {
    int key;
    Node *prev, *next;
    Node() { }
    Node(int k, Node* prev, Node* next)
        : key(k)
        , prev(prev)
        , next(next)
    {
    }
} Node;

// 双向链表
class DoubleList {
private:
    Node *begin, *end;

public:
    DoubleList();
    Node* insertNode(int);
    void deleteNode(Node*);
    Node* getLast();
    void moveToFirst(Node*);
};

class LRU {
private:
    int length, capacity;
    int count, missCount;
    std::unordered_map<int, Node*> cache;
    DoubleList d;

public:
    LRU(int);
    void get(int);
    float getMissingRate();
};