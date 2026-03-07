#ifndef MOMU_DATA_STRUCTURES_ARRAY_STACK_H
#define MOMU_DATA_STRUCTURES_ARRAY_STACK_H

#include "array_list.h"

namespace momu {
namespace data_structures {

template <typename T>
class ArrayStack {
   public:
    ArrayStack(size_t capacity = 2) : list(capacity) {}

    void Push(const T& data) { list.PushBack(data); }

    void Push(T&& data) { list.PushBack(std::move(data)); }

    template <typename... Args>
    void Emplace(Args&&... args) {
        list.EmplaceBack(std::forward<Args>(args)...);
    }

    void Pop() { list.PopBack(); }

    T& Top() { return list.Back(); }

    const T& Top() const { return list.Back(); }

    bool Empty() const { return list.Empty(); }

   private:
    ArrayList<T> list;
};

}  // namespace data_structures
}  // namespace momu

#endif  // MOMU_DATA_STRUCTURES_ARRAY_STACK_H