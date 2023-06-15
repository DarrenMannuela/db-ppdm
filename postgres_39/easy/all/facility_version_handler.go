package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_version

	for rows.Next() {
		var facility_version dto.Facility_version
		if err := rows.Scan(&facility_version.Facility_id, &facility_version.Facility_type, &facility_version.Source, &facility_version.Active_date, &facility_version.Active_ind, &facility_version.Constructed_date, &facility_version.Current_operator, &facility_version.Current_status_date, &facility_version.Description, &facility_version.Effective_date, &facility_version.Expiry_date, &facility_version.Facility_long_name, &facility_version.Facility_short_name, &facility_version.Facility_status_id, &facility_version.Inactive_date, &facility_version.Last_injection_date, &facility_version.Last_production_date, &facility_version.Last_reported_date, &facility_version.Latitude, &facility_version.Longitude, &facility_version.On_injection_date, &facility_version.On_production_date, &facility_version.Plot_name, &facility_version.Plot_symbol, &facility_version.Ppdm_guid, &facility_version.Remark, &facility_version.Status_type, &facility_version.Row_changed_by, &facility_version.Row_changed_date, &facility_version.Row_created_by, &facility_version.Row_created_date, &facility_version.Row_effective_date, &facility_version.Row_expiry_date, &facility_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_version to result
		result = append(result, facility_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityVersion(c *fiber.Ctx) error {
	var facility_version dto.Facility_version

	setDefaults(&facility_version)

	if err := json.Unmarshal(c.Body(), &facility_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	facility_version.Row_created_date = formatDate(facility_version.Row_created_date)
	facility_version.Row_changed_date = formatDate(facility_version.Row_changed_date)
	facility_version.Active_date = formatDateString(facility_version.Active_date)
	facility_version.Constructed_date = formatDateString(facility_version.Constructed_date)
	facility_version.Current_status_date = formatDateString(facility_version.Current_status_date)
	facility_version.Effective_date = formatDateString(facility_version.Effective_date)
	facility_version.Expiry_date = formatDateString(facility_version.Expiry_date)
	facility_version.Inactive_date = formatDateString(facility_version.Inactive_date)
	facility_version.Last_injection_date = formatDateString(facility_version.Last_injection_date)
	facility_version.Last_production_date = formatDateString(facility_version.Last_production_date)
	facility_version.Last_reported_date = formatDateString(facility_version.Last_reported_date)
	facility_version.On_injection_date = formatDateString(facility_version.On_injection_date)
	facility_version.On_production_date = formatDateString(facility_version.On_production_date)
	facility_version.Row_effective_date = formatDateString(facility_version.Row_effective_date)
	facility_version.Row_expiry_date = formatDateString(facility_version.Row_expiry_date)






	rows, err := stmt.Exec(facility_version.Facility_id, facility_version.Facility_type, facility_version.Source, facility_version.Active_date, facility_version.Active_ind, facility_version.Constructed_date, facility_version.Current_operator, facility_version.Current_status_date, facility_version.Description, facility_version.Effective_date, facility_version.Expiry_date, facility_version.Facility_long_name, facility_version.Facility_short_name, facility_version.Facility_status_id, facility_version.Inactive_date, facility_version.Last_injection_date, facility_version.Last_production_date, facility_version.Last_reported_date, facility_version.Latitude, facility_version.Longitude, facility_version.On_injection_date, facility_version.On_production_date, facility_version.Plot_name, facility_version.Plot_symbol, facility_version.Ppdm_guid, facility_version.Remark, facility_version.Status_type, facility_version.Row_changed_by, facility_version.Row_changed_date, facility_version.Row_created_by, facility_version.Row_created_date, facility_version.Row_effective_date, facility_version.Row_expiry_date, facility_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityVersion(c *fiber.Ctx) error {
	var facility_version dto.Facility_version

	setDefaults(&facility_version)

	if err := json.Unmarshal(c.Body(), &facility_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_version.Facility_id = id

    if facility_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_version set facility_type = :1, source = :2, active_date = :3, active_ind = :4, constructed_date = :5, current_operator = :6, current_status_date = :7, description = :8, effective_date = :9, expiry_date = :10, facility_long_name = :11, facility_short_name = :12, facility_status_id = :13, inactive_date = :14, last_injection_date = :15, last_production_date = :16, last_reported_date = :17, latitude = :18, longitude = :19, on_injection_date = :20, on_production_date = :21, plot_name = :22, plot_symbol = :23, ppdm_guid = :24, remark = :25, status_type = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where facility_id = :34")
	if err != nil {
		return err
	}

	facility_version.Row_changed_date = formatDate(facility_version.Row_changed_date)
	facility_version.Active_date = formatDateString(facility_version.Active_date)
	facility_version.Constructed_date = formatDateString(facility_version.Constructed_date)
	facility_version.Current_status_date = formatDateString(facility_version.Current_status_date)
	facility_version.Effective_date = formatDateString(facility_version.Effective_date)
	facility_version.Expiry_date = formatDateString(facility_version.Expiry_date)
	facility_version.Inactive_date = formatDateString(facility_version.Inactive_date)
	facility_version.Last_injection_date = formatDateString(facility_version.Last_injection_date)
	facility_version.Last_production_date = formatDateString(facility_version.Last_production_date)
	facility_version.Last_reported_date = formatDateString(facility_version.Last_reported_date)
	facility_version.On_injection_date = formatDateString(facility_version.On_injection_date)
	facility_version.On_production_date = formatDateString(facility_version.On_production_date)
	facility_version.Row_effective_date = formatDateString(facility_version.Row_effective_date)
	facility_version.Row_expiry_date = formatDateString(facility_version.Row_expiry_date)






	rows, err := stmt.Exec(facility_version.Facility_type, facility_version.Source, facility_version.Active_date, facility_version.Active_ind, facility_version.Constructed_date, facility_version.Current_operator, facility_version.Current_status_date, facility_version.Description, facility_version.Effective_date, facility_version.Expiry_date, facility_version.Facility_long_name, facility_version.Facility_short_name, facility_version.Facility_status_id, facility_version.Inactive_date, facility_version.Last_injection_date, facility_version.Last_production_date, facility_version.Last_reported_date, facility_version.Latitude, facility_version.Longitude, facility_version.On_injection_date, facility_version.On_production_date, facility_version.Plot_name, facility_version.Plot_symbol, facility_version.Ppdm_guid, facility_version.Remark, facility_version.Status_type, facility_version.Row_changed_by, facility_version.Row_changed_date, facility_version.Row_created_by, facility_version.Row_effective_date, facility_version.Row_expiry_date, facility_version.Row_quality, facility_version.Facility_id)
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

func PatchFacilityVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_version set "
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
		} else if key == "active_date" || key == "constructed_date" || key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "inactive_date" || key == "last_injection_date" || key == "last_production_date" || key == "last_reported_date" || key == "on_injection_date" || key == "on_production_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where facility_id = :id"

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

func DeleteFacilityVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_version dto.Facility_version
	facility_version.Facility_id = id

	stmt, err := db.Prepare("delete from facility_version where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_version.Facility_id)
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


