import tutorial_pb2

# create a new Person object
person = tutorial_pb2.Person()
person.name = "John Doe"
person.id = 1234
person.email = "johndoe@example.com"

# add a phone number to the Person object
phone = person.phones.add()
phone.number = "555-1234"
phone.type = tutorial_pb2.Person.PhoneNumber.WORK

# read existing AddressBook from file or create a new one
address_book = tutorial_pb2.AddressBook()
try:
    with open("address_book.bin", "rb") as f:
        address_book.ParseFromString(f.read())
except IOError:
    print("Creating a new address book.")

# add the new Person object to the AddressBook object
address_book.people.append(person)

# write the updated AddressBook object to file
with open("address_book.bin", "wb") as f:
    f.write(address_book.SerializeToString())

# read the updated AddressBook object from file
with open("address_book.bin", "rb") as f:
    address_book.ParseFromString(f.read())

# print the updated AddressBook object
print(address_book)
