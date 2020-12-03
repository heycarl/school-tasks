program task1_11V_home;

var A: array[1..15] of integer;
var sum: integer;

begin

for var i:=1 to 15 do begin
  writeln('Enter ' + i + ' num: ');
  read(A[i]);
end;

for var i:=1 to 15 do begin
  sum := sum + A[i];
end;

writeln('Среднее арифметическое: ', sum / 15:0:2);

end.