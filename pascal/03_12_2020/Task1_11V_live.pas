program task1_11V_live;

var A: array[1..2,1..5] of integer;

begin


for var i:=1 to 2 do begin
  for var j:=1 to 5 do
  begin
     read(A[i,j]);
  end;
end;

for var i:=1 to 2 do begin
  writeln();
  for var j:=1 to 5 do
  begin
     write(A[i,j], ' ');
  end;
end;

end.