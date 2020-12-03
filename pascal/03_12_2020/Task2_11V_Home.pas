program task2_11V_home;

var A: array[1..7] of real;
var sum_otric: real;
var multi_poloz: real;

begin

multi_poloz := 1;

for var i:=1 to 7 do begin
  writeln('Enter ' + i + ' num: ');
  read(A[i]);
end;

for var i:=1 to 7 do begin
  if A[i] > 0 then
    multi_poloz := multi_poloz *A[i] ;
  if A[i] < 0 then
    sum_otric := sum_otric  + A[i] ;
end;

writeln('Сумма отрицательных: ', sum_otric:0:2);
writeln('Произведение положительных: ', multi_poloz:0:2);

end.