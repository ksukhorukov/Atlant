#!/usr/bin/env ruby

if ARGV.count != 1
	puts "\nGenerates CSV files for testing purposes\n"
	puts "\nUsage: #{$0} custome_file_name.csv"
	puts "\nExample: #{$0} sample_set_1.csv\n\n"
	exit
end

file_name = ARGV[0]

fp = File.open(file_name, 'w')

fp.print("PRODUCT NAME;PRICE\n")

1000.times do |n|
	product = "test_product_#{rand(1000) + 1}"
	price = (rand(1000) + 1) / 100.0
	fp.print("#{product};#{price}\n")
end

fp.close

puts "[+] Done!"