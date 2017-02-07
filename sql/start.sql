SELECT * FROM start.FileInfo;
SELECT FID,FileName,Extname,FileSize FROM FileInfo where FID=(SELECT MIN(FID) FROM FileInfo WHERE IsTranscode = 0);