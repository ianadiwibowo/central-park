# frozen_string_literal: true

# Complete the commonChild function below.
def common_child(s1, s2)
  s1_a = s1.chars
  s2_a = s2.chars

  # Remove non-shared characters
  intersection = s1_a & s2_a
  s1_clean = s1_a.select { |c| intersection.include?(c) }
  s2_clean = s2_a.select { |c| intersection.include?(c) }
  return 0 if s1_clean.empty? || s2_clean.empty?

  # Find longest shared combination length
  l = s1_clean.length < s2_clean.length ? s1_clean.length : s2_clean.length
  l.downto(1).each do |i|
    s1_comb = s1_clean.combination(i).to_a
    counter = {}

    (0..s1_comb.length - 1).each do |j|
      counter[s1_comb[j].join] = counter[s1_comb[j].join].to_i + 1
    end

    s2_comb = s2_clean.combination(i).to_a
    (0..s2_comb.length - 1).each do |j|
      next if counter[s2_comb[j].join].nil?

      return i
    end
  end

  0
end

# Unit tests
s1_cases = %w[
  ABCD
  HARRY
  AA
  SHINCHAN
  ABCDEF
]
s2_cases = %w[
  ABDC
  SALLY
  BB
  NOHARAAA
  FBDAMN
]
expectations = [
  3,
  2,
  0,
  3,
  2
]
s1_cases.each_with_index do |v, i|
  result = common_child(v, s2_cases[i])
  puts "Passed: #{result == expectations[i]}." \
       "Test: #{v}, #{s2_cases[i]}." \
       "Expected: #{expectations[i]}." \
       "Got: #{result}."
end
