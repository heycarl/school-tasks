program task1;

var A: array[1..7] of integer;

begin

for var i:=1 to 7 do begin
  writeln('Enter ', i, ' value: ');
  read(A[i]);
end;

for var i:=1 to 7 do 
  write(A[i], ' ');

end.