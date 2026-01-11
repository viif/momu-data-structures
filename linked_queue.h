#ifndef MOMU_DATA_STRUCTURES_LINKED_QUEUE_H
#define MOMU_DATA_STRUCTURES_LINKED_QUEUE_H

#include "linked_list.h"

namespace momu {
namespace data_structures {

template <typename T>
class LinkedQueue {
   public:
    void Push(const T& data) { list.PushBack(data); }

    void Push(T&& data) { list.PushBack(std::move(data)); }

    template <typename... Args>
    void Emplace(Args&&... args) {
        list.EmplaceBack(std::forward<Args>(args)...);
    }

    void Pop() { list.PopFront(); }

    T& Front() { return list.Front(); }

    const T& Front() const { return list.Front(); }

    T& Back() { return list.Back(); }

    const T& Back() const { return list.Back(); }

    bool Empty() const { return list.Empty(); }

   private:
    LinkedList<T> list;
};

}  // namespace data_structures
}  // namespace momu

#endif  // MOMU_DATA_STRUCTURES_LINKED_QUEUE_H