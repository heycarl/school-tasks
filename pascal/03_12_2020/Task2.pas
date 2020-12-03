program task2;

var A: array[1..10] of integer;
var n: integer;

begin

for var i:=1 to 10 do begin
  writeln('Enter temperature in ' + i + ' day: ');
  read(A[i]);
end;

for var i:=1 to 10 do begin
  if A[i] < 5 then
    n := n +1 ;
end;

writeln('There was ', n, ' days with temp < 5');

end.