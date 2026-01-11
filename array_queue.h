#ifndef MOMU_DATA_STRUCTURES_ARRAY_QUEUE_H
#define MOMU_DATA_STRUCTURES_ARRAY_QUEUE_H

#include "array_deque.h"

namespace momu {
namespace data_structures {

template <typename T>
class ArrayQueue {
   public:
    ArrayQueue(size_t capacity = 2) : deque_(capacity) {}

    void Push(const T& value) { deque_.PushBack(value); }

    void Push(T&& value) { deque_.PushBack(std::move(value)); }

    template <typename... Args>
    T& Emplace(Args&&... args) {
        return deque_.EmplaceBack(std::forward<Args>(args)...);
    }

    void Pop() { deque_.PopFront(); }

    T& Front() { return deque_.Front(); }

    const T& Front() const { return deque_.Front(); }

    T& Back() { return deque_.Back(); }

    const T& Back() const { return deque_.Back(); }

    bool Empty() const { return deque_.Empty(); }

   private:
    ArrayDeque<T> deque_;
};

}  // namespace data_structures
}  // namespace momu

#endif  // MOMU_DATA_STRUCTURES_ARRAY_QUEUE_H