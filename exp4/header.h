#include <queue>
#include <unordered_map>

class PermutationAlgo {
protected:
    int capacity;
    int count, missCount;

public:
    virtual void get(int) = 0;
    float getMissingRate();
    int getCount();
    int getMissCount();
};

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

class LRU : public PermutationAlgo {
private:
    std::unordered_map<int, Node*> cache;
    DoubleList d;

public:
    LRU(int);
    void get(int);
};

class FIFO : public PermutationAlgo {
private:
    std::unordered_map<int, bool> cache;
    std::queue<int> queue;

public:
    FIFO(int);
    void get(int);
};