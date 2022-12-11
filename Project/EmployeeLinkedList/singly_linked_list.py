import datetime

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class LinkedList:
    def __init__(self):
        self.head = None
        self.tail = None

    def AddNode(self, data):
        temp = Node(data)
        if self.head == None:
            self.head = temp
            self.tail = temp
        else:
            self.tail.next = temp
            self.tail = temp

    def RemoveNode(self, index):
        temp = self.head
        if temp.data == index:
            temp = self.head = temp.next
        while ( temp.next != None):
            if temp.next.data == index:
                node = temp.next
                temp.next  = temp.next.next
                del node
            else:
                temp = temp.next

    def RemoveNodeAt(self, condition):
        temp = self.head
        if condition(temp.data):
            temp = self.head = temp.next
        while ( temp.next != None):
            if condition(temp.data):
                node = temp.next
                temp.next  = temp.next.next
                del node
            else:
                temp = temp.next

    def Print(self):
        temp =  self.head
        count = 1
        while ( temp.next != None):
            print(temp.data)
            temp = temp.next
            if temp.next ==  None:
                print(temp.data)
            count += 1

def get_input_until(ask, condition):
    while True:
        result = input(ask)
        if condition(result):
            break
    return result

def is_valid_date(datestring):
    try:
        datetime.datetime.strptime(datestring, '%Y-%m-%d')
        return True
    except ValueError:
        return False

if __name__ == '__main__':
    employees = LinkedList()
    menu = -1
    while menu != 5:
        print("========================= PROGRAM DATABASE KARYAWAN =========================")
        print("1. Menambah data karyawan")
        print("2. Menampilkan semua daftar karyawan")
        print("3. Hapus data karyawan berdasarkan ID")
        print("4. Hapus semua data karyawan")
        print("5. Keluar program")
        menu = int(input("Masukkan nomor yang anda pilih: "))
        if menu == 1:
            employee = {}
            employee["ID"] = int(get_input_until("Masukkan ID karyawan: ", lambda x : int(x) <= 99999))
            employee["Name"] = get_input_until("Masukkan nama lengkap karyawan: ", lambda x : len(x) <= 30)
            employee["Birthdate"] = get_input_until("Masukkan tanggal lahir karyawan (ex. 2003-12-12): ", is_valid_date)
            employee["Birthplace"] = get_input_until("Masukkan tempat lahir karyawan: ", lambda x: len(x) <= 30)
            employees.AddNode(employee)
        elif menu == 2:
            employees.Print()
        elif menu == 3:
            employeeId = int(get_input_until("Masukkan ID karyawan: ", lambda x : int(x) <= 99999))
            employees.RemoveNodeAt(lambda x: x["ID"] == employeeId)
        elif menu == 4:
            employees = LinkedList()
        else:
            continue
