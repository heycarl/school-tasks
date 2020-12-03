program task2_11V_live;

var A: array[1..3,1..2] of integer;

begin


for var i:=1 to 3 do begin
  for var j:=1 to 2 do
  begin
     A[i,j] := i+j;
  end;
end;

for var i:=1 to 3 do begin
  writeln();
  for var j:=1 to 2 do
  begin
     write(A[i,j], ' ');
  end;
end;

end.