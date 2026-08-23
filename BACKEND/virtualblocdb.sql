create database virtualbloc;
use virtualbloc;

CREATE TABLE icono (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    frames JSON NOT NULL
);

CREATE TABLE libro (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    color CHAR(6) NOT NULL,
    icono INT NOT NULL,
    hojas INT DEFAULT 0,
    FOREIGN KEY (icono) REFERENCES icono(id)
);

CREATE TABLE hoja (
    id INT AUTO_INCREMENT PRIMARY KEY,
    libro INT NOT NULL,
    numero INT NOT NULL,
    texto VARCHAR(9000) NOT NULL,
    FOREIGN KEY (libro) REFERENCES libro(id)
);

DELIMITER //
CREATE PROCEDURE GetLibros()
BEGIN
	select 
		l.id, l.nombre, l.color, i.nombre as icono,l.hojas 
	from libro as l 
	inner join icono as i 
		on i.id = l.icono;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE SaveLibro(IN pid INT,IN pnombre VARCHAR(100),IN pcolor CHAR(6),IN picono VARCHAR(100),IN phojas INT)
BEGIN
	DECLARE idicono INT;
    SELECT id from icono where nombre=picono into idicono;
	IF pid < 0 THEN
		insert into libro values (null,pnombre,pcolor,idicono,phojas);
	ELSE
		update libro set nombre=pnombre,color=pcolor,icono=idicono,hojas=phojas where id=pid;
	END IF;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE SaveHojas(IN plibro INT,IN pnumero INT,IN ptexto VARCHAR(9000))
BEGIN
	DECLARE numhojas INT;
    SELECT hojas from libro where id=plibro into numhojas;
	IF pnumero > numhojas THEN
		insert into hoja values (null,plibro,pnumero,ptexto);
	ELSE
		update hoja set texto=ptexto where numero=pnumero and libro=plibro;
	END IF;
END //
DELIMITER ;

call SaveLibro(3,'wa','che','A',1);
select * from icono
