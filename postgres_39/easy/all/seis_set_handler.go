package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_set

	for rows.Next() {
		var seis_set dto.Seis_set
		if err := rows.Scan(&seis_set.Seis_set_subtype, &seis_set.Seis_set_id, &seis_set.Acqtn_design_id, &seis_set.Active_ind, &seis_set.Area_id, &seis_set.Area_size, &seis_set.Area_size_ouom, &seis_set.Area_type, &seis_set.Coord_acquisition_id, &seis_set.Current_seis_status, &seis_set.Effective_date, &seis_set.End_date, &seis_set.Expiry_date, &seis_set.Finance_id, &seis_set.Geographic_coord_system_id, &seis_set.Local_coord_system_id, &seis_set.Max_latitude, &seis_set.Max_longitude, &seis_set.Min_latitude, &seis_set.Min_longitude, &seis_set.Ppdm_guid, &seis_set.Preferred_name, &seis_set.Remark, &seis_set.Source, &seis_set.Start_date, &seis_set.Xy_coord_system_id, &seis_set.Row_changed_by, &seis_set.Row_changed_date, &seis_set.Row_created_by, &seis_set.Row_created_date, &seis_set.Row_effective_date, &seis_set.Row_expiry_date, &seis_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_set to result
		result = append(result, seis_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSet(c *fiber.Ctx) error {
	var seis_set dto.Seis_set

	setDefaults(&seis_set)

	if err := json.Unmarshal(c.Body(), &seis_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	seis_set.Row_created_date = formatDate(seis_set.Row_created_date)
	seis_set.Row_changed_date = formatDate(seis_set.Row_changed_date)
	seis_set.Effective_date = formatDateString(seis_set.Effective_date)
	seis_set.End_date = formatDateString(seis_set.End_date)
	seis_set.Expiry_date = formatDateString(seis_set.Expiry_date)
	seis_set.Start_date = formatDateString(seis_set.Start_date)
	seis_set.Row_effective_date = formatDateString(seis_set.Row_effective_date)
	seis_set.Row_expiry_date = formatDateString(seis_set.Row_expiry_date)






	rows, err := stmt.Exec(seis_set.Seis_set_subtype, seis_set.Seis_set_id, seis_set.Acqtn_design_id, seis_set.Active_ind, seis_set.Area_id, seis_set.Area_size, seis_set.Area_size_ouom, seis_set.Area_type, seis_set.Coord_acquisition_id, seis_set.Current_seis_status, seis_set.Effective_date, seis_set.End_date, seis_set.Expiry_date, seis_set.Finance_id, seis_set.Geographic_coord_system_id, seis_set.Local_coord_system_id, seis_set.Max_latitude, seis_set.Max_longitude, seis_set.Min_latitude, seis_set.Min_longitude, seis_set.Ppdm_guid, seis_set.Preferred_name, seis_set.Remark, seis_set.Source, seis_set.Start_date, seis_set.Xy_coord_system_id, seis_set.Row_changed_by, seis_set.Row_changed_date, seis_set.Row_created_by, seis_set.Row_created_date, seis_set.Row_effective_date, seis_set.Row_expiry_date, seis_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSet(c *fiber.Ctx) error {
	var seis_set dto.Seis_set

	setDefaults(&seis_set)

	if err := json.Unmarshal(c.Body(), &seis_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_set.Seis_set_subtype = id

    if seis_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_set set seis_set_id = :1, acqtn_design_id = :2, active_ind = :3, area_id = :4, area_size = :5, area_size_ouom = :6, area_type = :7, coord_acquisition_id = :8, current_seis_status = :9, effective_date = :10, end_date = :11, expiry_date = :12, finance_id = :13, geographic_coord_system_id = :14, local_coord_system_id = :15, max_latitude = :16, max_longitude = :17, min_latitude = :18, min_longitude = :19, ppdm_guid = :20, preferred_name = :21, remark = :22, source = :23, start_date = :24, xy_coord_system_id = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where seis_set_subtype = :33")
	if err != nil {
		return err
	}

	seis_set.Row_changed_date = formatDate(seis_set.Row_changed_date)
	seis_set.Effective_date = formatDateString(seis_set.Effective_date)
	seis_set.End_date = formatDateString(seis_set.End_date)
	seis_set.Expiry_date = formatDateString(seis_set.Expiry_date)
	seis_set.Start_date = formatDateString(seis_set.Start_date)
	seis_set.Row_effective_date = formatDateString(seis_set.Row_effective_date)
	seis_set.Row_expiry_date = formatDateString(seis_set.Row_expiry_date)






	rows, err := stmt.Exec(seis_set.Seis_set_id, seis_set.Acqtn_design_id, seis_set.Active_ind, seis_set.Area_id, seis_set.Area_size, seis_set.Area_size_ouom, seis_set.Area_type, seis_set.Coord_acquisition_id, seis_set.Current_seis_status, seis_set.Effective_date, seis_set.End_date, seis_set.Expiry_date, seis_set.Finance_id, seis_set.Geographic_coord_system_id, seis_set.Local_coord_system_id, seis_set.Max_latitude, seis_set.Max_longitude, seis_set.Min_latitude, seis_set.Min_longitude, seis_set.Ppdm_guid, seis_set.Preferred_name, seis_set.Remark, seis_set.Source, seis_set.Start_date, seis_set.Xy_coord_system_id, seis_set.Row_changed_by, seis_set.Row_changed_date, seis_set.Row_created_by, seis_set.Row_effective_date, seis_set.Row_expiry_date, seis_set.Row_quality, seis_set.Seis_set_subtype)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchSeisSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_set set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_set_subtype = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteSeisSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_set dto.Seis_set
	seis_set.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_set where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_set.Seis_set_subtype)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


