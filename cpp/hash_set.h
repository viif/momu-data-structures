#ifndef MOMU_DATA_STRUCTURES_HASH_SET_H
#define MOMU_DATA_STRUCTURES_HASH_SET_H

#include "linear_probing_hash_map.h"

namespace momu {
namespace data_structures {

template <typename K, class Hash>
class HashSet {
   public:
    // 增
    bool Insert(const K& key) {
        if (map_.Contains(key)) {
            return false;
        } else {
            map_[key] = kValue_;
            return true;
        }
    }

    bool Insert(K&& key) {
        if (map_.Contains(std::move(key))) {
            return false;
        } else {
            map_[std::move(key)] = kValue_;
            return true;
        }
    }

    // 删
    bool Erase(const K& key) { return map_.Erase(key); }

    bool Erase(K&& key) { return map_.Erase(std::move(key)); }

    void Clear() { map_.Clear(); }

    // 查
    bool Contains(const K& key) { return map_.Contains(key); }

    bool Contains(K&& key) { return map_.Contains(std::move(key)); }

    // 工具函数
    size_t Size() const { return map_.Size(); }

    bool Empty() const { return map_.Empty(); }

   private:
    LinearProbingHashMap<K, char, Hash> map_;
    const char kValue_{'0'};
};

}  // namespace data_structures
}  // namespace momu

#endif  // MOMU_DATA_STRUCTURES_HASH_SET_H