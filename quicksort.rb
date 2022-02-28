require 'stringio'

def quicksort(values, low, high)
  if low < high
    if high - low <= 10
      insertionsort(values, low, high)
    else
      pivot = partition(values, low, high)
      quicksort(values, low, pivot)
      quicksort(values, pivot + 1, high)
    end
  end
end

def partition(values, low, high)
  pivot = values[(low + high) / 2]
  left = low - 1
  right = high + 1

  loop do
    loop do
      left += 1
      break if values[left] >= pivot
    end

    loop do
      right -= 1
      break if values[right] <= pivot
    end

    return right if left >= right

    values[left], values[right] = values[right], values[left]
  end
end

def insertionsort(values, low, high)
  ((low + 1)..high).each do |i|
    value = values[i]
    j = i

    while j > low && values[j - 1] > value
      values[j] = values[j - 1]
      j -= 1
    end

    values[j] = value
  end
end

inp = (gets.split ' ').map(&:to_i).drop(1)

quicksort(inp, 0, inp.length - 1)

output = StringIO.new
inp.each do |val|
  output << val
  output << ' '
end

print(output.string)
