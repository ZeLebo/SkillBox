% To be shorted the following code works in this way:

% Firstly, we will build up all the possible exam combinations and free rooms
% We need to such exams, that wasnt in the schedule
% Then we calculate the fines the variant
% Produce this till all the exams are not done

% TODO more comments, maybe in javadoc style

%:- use_module(small_data).
:- use_module(large_data).
:- use_module(utils).
:- use_module(cost).



% Checks whether the day is in session period
fitsTheSeason(Day):- 
    ex_season_starts(StartDay), % get the start dat of session
    ex_season_ends(EndDay), % get the end day of session
    between(StartDay, EndDay, Day). % check that the day is in the season


% check whether the room is avaliable for all exam duration
fitsTheTime(ExamID, RoomID, Day, ExamStart) :- 
    exam_duration(ExamID, Duration), % getting the duration of the exam
    classroom_available(RoomID, Day, From, Till), % check whether the room is avaliable
    ExamStart >= From, NewTime is ExamStart + Duration, NewTime =< Till. % check the time

% If students can fit the capacity of the room
fitStudentsAmount(ExamID, RoomID) :-
    st_group(ExamID, Students), % get the list of students to pass the exam
    length(Students, StudentCount), % number of students to pass the exam
    classroom_capacity(RoomID, RoomCapacity),  % getting the capacity of room
    StudentCount =< RoomCapacity. % checking that room can accommodate all the students

% Check that specifiend exam are not colliding
timeCollision(ExamID0, ExamID, From0, From):-
    exam_duration(ExamID0, D0), % getting the duration of first exam
    exam_duration(ExamID, D), % getting the duration of second exam
    E0 is From0 + D0, E is From + D, % calculate the time of exams duration
    (between(From0, E0, From); between(From, E, From0)). % check that they don't collide


% Check that two exams are not in the same room
collision(ExamID0, RID0, Day0, From0, [event(ExamID, RoomID, Day, From)|_]) :-
    RID0 == RoomID, % check the rooms are the same
    Day0 == Day, % check the days are the same
    timeCollision(ExamID0, ExamID, From0, From). % check that exams don't collide'

% Checks whether the students have to go to multyple exams at one time
collision(ExamID0, _, Day0, From0, [event(ExamID, _, Day, From)|_]):-
    Day0 == Day, % check that the days are same
    student_follows_both_classes(ExamID0, ExamID), % check that the students have to go to multyple exams
    timeCollision(ExamID0, ExamID, From0, From). % check the exams don't collide

% Checks whether the teacher has to go to multyple exams at one time
collision(ExamID0, _, Day0, From0, [event(ExamID, _, Day, From)| _]) :-
    Day0 == Day, % check that the days are the same
    teacher_teaches_both_classes(ExamID0, ExamID), % check that teacher has two exams to get
    timeCollision(ExamID0, ExamID, From0, From). % check that exams don't collide

% Checking the list of events recursively
collision(ExamID0, RID0, Day0, From0, [_| Others]) :- collision(ExamID0, RID0, Day0, From0, Others).


% Could become event, if all of the above is done
% Fits the time, the capacity, no double exam
% isEvent is True if all the checks are done
% Event day is in season interval
% Room capacity fits the group and  is avaliable for all exam duration
% It's the first appearing of the exam in the schedule
% The exam has no collision for time
isEvent(CurState, ExamID, RoomID, Day, From) :-
    not(member(event(ExamID, _, _, _), CurState)), % check that exam is not in the schedule
    fitsTheSeason(Day), % check that day is in the session period
    fitsTheTime(ExamID, RoomID, Day, From), % check that time fits the available time of room
    fitStudentsAmount(ExamID, RoomID), % check that room fits students amount
    not(collision(ExamID, RoomID, Day, From, CurState)). % no collision for timing

% Find all possible options for the start of the exam
cycleForHours(_, _, _ , _, From, Till, Container, Container) :- From > Till + 1.
cycleForHours(CurState, ExamID, RoomID, Day, CurStartTime, Till, Container, Result):-
    not(isEvent(CurState, ExamID, RoomID, Day, CurStartTime)), !, % if the hours are not correct
    NewStartTime is CurStartTime + 1, % increase the current start time
    cycleForHours(CurState, ExamID, RoomID, Day, NewStartTime, Till, Container, Result). % call recursively for rest hours
cycleForHours(CurState, ExamID, RoomID, Day, CurStartTime, Till, Container, Result):-
    isEvent(CurState, ExamID, RoomID, Day, CurStartTime), % if the hours are correct
    append(Container, [event(ExamID, RoomID, Day, CurStartTime)], NewContainer), !, % add the start hour to container
    NewStartTime is CurStartTime + 1, % increase the current start time
    cycleForHours(CurState, ExamID, RoomID, Day, NewStartTime, Till, NewContainer, Result). % call recursively for the rest hours


% Find all possible options for the next exam
cycleForRooms(_, _, [], Container, Container).
cycleForRooms(CurState, ExamID, [(RoomID, Day, From, Till) | Others], Container, Result) :-
    exam_duration(ExamID, Duration), ExamTill is Till - Duration, % get exam duration
    cycleForHours(CurState, ExamID, RoomID, Day, From, ExamTill, [], Events), % seach the start hour for the exam
    append(Container, Events, NewContainer), !, % add the rooms to Container
    cycleForRooms(CurState, ExamID, Others, NewContainer, Result). % call recursively for rest rooms

% Go throught the options of the next exam
cycleForExams(_, [], _, Container, Container).
cycleForExams(CurState, [ExamID|Others], Rooms, Container, Result):-
    !, cycleForRooms(CurState, ExamID, Rooms, [], Events), % search the room for the next exam
    append(Container, Events, NewContainer), % add new event to container
    cycleForExams(CurState, Others, Rooms, NewContainer, Result). % call for the rest exams

% Make the list of the events (Result)
% ExamIDS - list of exams
% Rooms - list with nodes RoomID, day, from, till
getEvents((State, _, _), ExamIDS, Rooms, Result) :- !, cycleForExams(State, ExamIDS, Rooms, [], Result).


% Make nodes from events
% Contatiner to store the result
makeNodesFromEvents([], _, Container, Container).
makeNodesFromEvents([Event|Others], (State, Length, H), Container, Result) :-
    append(State, [Event], StateNext), % construct new state from the previous and event
    NewL is Length + 1, % increase the lenth of schedule
    cost(schedule(StateNext), NewCost), % recalculate cost for the new schedule
    append(Container, [(StateNext, NewL, NewCost)], Nodes), !, % add to accumulator the new node
    makeNodesFromEvents(Others, (State, Length, H), Nodes, Result). % recursively calling


% Check whether the exam was in the schedule
isInTable((State, _, _), ExamID) :- member(event(ExamID, _, _, _), State).
% Check whether this exam wasn't' in current schedule
isNotUsedExam(ExamID, CurNode) :- exam(ExamID, _), not(isInTable(CurNode, ExamID)).
% Checks whether the room is vacant (avaliable to use in the schedule)
isFreeRoom((RoomID, Day, From, Till)) :- classroom_available(RoomID, Day, From, Till).


% Function to get all of the neighbours of current node
% CurNode - current node
% Neigbours - list of neighbours of current node
getNeighbours(CurNode, Neigbours):-
    findall(ExamID, isNotUsedExam(ExamID, CurNode), ExamIDS), % exams to proofcheck
    findall(Room, isFreeRoom(Room), Rooms), % find the rooms to check
    getEvents(CurNode, ExamIDS, Rooms, Events), % collect the list of variants for the next exam
    makeNodesFromEvents(Events, CurNode, [], Nodes), 
    Neigbours = Nodes, !. % make nodes from events


% This function is printing the data
% Represented in the schedule
dataPrint([]).
dataPrint([event(ExamID, RoomID, Day, Hour)|Tail]) :-
    exam(ExamID, EName), nl,
    write("Exam: "), writeln(EName),
    write("Room: "), writeln(RoomID),
    write("Date: "), writeln(Day),
    write("Starts at: "), write(Hour), writeln(" hour."),
    exam_duration(ExamID, Duration),
    Till is Hour + Duration,
    write("Ends at: "), write(Till), writeln(" hour."), !, dataPrint(Tail).

% Print the header of schedule (exam ammount + fines)
schedulePrint(State, Length, H) :-
    write("The schedule length: "), writeln(Length),
    write("All the fines: "), writeln(H),
    dataPrint(State).

% get the minimal (first) element from list
getMinimal([Head|_], First) :- First = Head, !.

% Get the neighbours of the node
mainAlgoCycle((State, Length, Fines), Nodes) :- 
    Length == Nodes, % if all the exams are "used"
    schedulePrint(State, Length, Fines), !. % print the result

mainAlgoCycle((State, Length, Fines), Nodes) :- Length < Nodes,
    getNeighbours((State, Length, Fines), Neighbours), % all the possible variants of the next exam
    sort(2, =<, Neighbours, Sorted), % get the minimal for fine exam
    getMinimal(Sorted, Cheapest), % pick it up
    mainAlgoCycle(Cheapest, Nodes), !. % recursively calling


% The main function, starts automaticaly
% Start point - list of parameters is empty
main :-
    findall(ExamID, exam(ExamID, _), ExamIDs), % collecting all exams to list
    prepare_env(ExamIDs), % preparing environment
    length(ExamIDs, N), % number of exams
    mainAlgoCycle(([], 0, 0), N), % main algorithm
    halt. % exit the programm

% called when starts from command line
:- initialization main.