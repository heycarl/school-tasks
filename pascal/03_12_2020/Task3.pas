program task3;

var n_elements: integer;
var chet_otrit: integer;
var nechet_poloz: integer;
var fst_lst_sum: real;
var A: array[1..99] of real;

begin

writeln('Enter n elements in array: ');
read(n_elements);

for var i:=1 to n_elements do begin
  A[i] := sqrt(i+5) / (8 - sqr(i))
end;

for var i:=1 to n_elements do begin
  if i mod 2 = 0 then begin
    if A[i] < 0 then
      chet_otrit := chet_otrit + 1;
  end;
  if i mod 2 <> 0 then begin
    if A[i] > 0 then
      nechet_poloz := chet_otrit + 1;
  end;
end;

fst_lst_sum := A[1] + A[n_elements];

writeln('Четных отрицательных: ', chet_otrit);
writeln('Нечетных положительных: ', nechet_poloz);
writeln('Сумма первого и последнего: ', fst_lst_sum:0:2);

end.