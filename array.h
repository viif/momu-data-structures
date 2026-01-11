#ifndef MOMU_DATA_STRUCTURES_ARRAY_H
#define MOMU_DATA_STRUCTURES_ARRAY_H

namespace momu {
namespace data_structures {

template <typename T, size_t S>
class Array {
   public:
    constexpr size_t Size() const { return S; }

    T& operator[](size_t index) { return data_[index]; }
    const T& operator[](size_t index) const { return data_[index]; }

    T* Data() { return data_; }
    const T* Data() const { return data_; }

   private:
    T data_[S];
};

}  // namespace data_structures
}  // namespace momu

#endif  // MOMU_DATA_STRUCTURES_ARRAY_H