def countStudents(students= list, sandwiches= list):
        
        hs = len(students)
        eaten = []
        for i in range(len(sandwiches)):
            if students[i] == sandwiches[i]:
                eaten.append(students[i])
                print('removing')
            else:
                students.append(students.pop(students.index(students[i])))

        swce = hs - (len(eaten))
        return swce

students = [0,0,1,1,0,1,0]
meals = [0,0,0,1,0,1,1]

new_students = countStudents(students, meals)
print(new_students)


#wrong we're not keeping count of whether the student has looked at a sandwich twice. We need a counter 
#which will determine when everyone has tried to eat such as 

"""
check counter against length of students 
if a student eats remove from list
if they dont eat return to back of queue and start counter 
stop searching once counter is greater than amount of students in queue"""