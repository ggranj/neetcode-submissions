type DynamicArray struct {
    size int
    cap int
    list []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        size: 0,
        cap: capacity,
        list: make([]int, capacity), // [0] // len 1 cap 1
        }
}

func (da *DynamicArray) Get(i int) int {
    return da.list[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.list[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.size >= da.cap { // 0 >= 1
        da.resize()
        da.list[da.size] = n
        da.size++
    } else {
        da.list[da.size] = n // da.list[0]
        da.size++
    }
}

func (da *DynamicArray) Popback() int {
    last := da.list[da.size-1]
    da.size = da.size - 1

    return last
}

func (da *DynamicArray) resize() {
    doubled := make([]int, len(da.list)*2)
    copy(doubled, da.list)

    da.list = doubled
    da.cap = len(da.list)
}

func (da *DynamicArray) GetSize() int {
    return da.size
}

func (da *DynamicArray) GetCapacity() int {
    return da.cap
}
