#ifndef MOMU_DATA_STRUCTURES_LINKED_STACK_H
#define MOMU_DATA_STRUCTURES_LINKED_STACK_H

#include "linked_list.h"

namespace momu {
namespace data_structures {

template <typename T>
class LinkedStack {
   public:
    void Push(const T& data) { list.PushFront(data); }

    void Push(T&& data) { list.PushFront(std::move(data)); }

    template <typename... Args>
    void Emplace(Args&&... args) {
        list.EmplaceFront(std::forward<Args>(args)...);
    }

    void Pop() { list.PopFront(); }

    T& Top() { return list.Front(); }

    const T& Top() const { return list.Front(); }

    bool Empty() const { return list.Empty(); }

   private:
    LinkedList<T> list;
};

}  // namespace data_structures
}  // namespace momu

#endif  // MOMU_DATA_STRUCTURES_LINKED_STACK_H