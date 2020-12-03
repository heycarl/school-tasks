program task3thru_11V_live;

var A: array[1..3,1..5] of integer;
var sum_otric:  integer;
var n_chet: integer;
var zero_one_two:  integer;
begin


for var i:=1 to 3 do begin
  for var j:=1 to 5 do
  begin
     read(A[i,j]);
  end;
end;

for var i:=1 to 3 do begin
  writeln();
  for var j:=1 to 5 do
  begin
    if A[i,j] < 0 then
      sum_otric := sum_otric + A[i,j];
    if j mod 2 = 0 then
      n_chet := n_chet + 1;
    if A[i,j] = 0 then
      zero_one_two := zero_one_two + 1;
    if A[i,j] = 1 then
      zero_one_two := zero_one_two + 1;
    if A[i,j] = 2 then
      zero_one_two := zero_one_two + 1;
     write(A[i,j], ' ');
  end;
end;

writeln();
writeln('Сумма отрицательных элементов: ', sum_otric);
writeln('Количество элементов, стоящих на четных местах: ', n_chet);
writeln('Количество элментов, равных 0, 1 или 2: ', zero_one_two);

end.