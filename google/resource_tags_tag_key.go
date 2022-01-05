// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
  "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)



    
func resourceTagsTagKey() *schema.Resource {
    return &schema.Resource{
        Create: resourceTagsTagKeyCreate,
        Read: resourceTagsTagKeyRead,
        Update: resourceTagsTagKeyUpdate,
        Delete: resourceTagsTagKeyDelete,

        Importer: &schema.ResourceImporter{
            State: resourceTagsTagKeyImport,
        },

        Timeouts: &schema.ResourceTimeout {
            Create: schema.DefaultTimeout(4 * time.Minute),
            Update: schema.DefaultTimeout(4 * time.Minute),
            Delete: schema.DefaultTimeout(4 * time.Minute),
        },



        Schema: map[string]*schema.Schema{
"parent": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
	Description: `Input only. The resource name of the new TagKey's parent. Must be of the form organizations/{org_id}.`,
},
"short_name": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
		ValidateFunc: validation.StringLenBetween(1, 63),
		Description: `Input only. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace.

The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.`,
},
"description": {
    Type: schema.TypeString,
    Optional: true,
		ValidateFunc: validation.StringLenBetween(0, 256),
		Description: `User-assigned description of the TagKey. Must not exceed 256 characters.`,
},
"create_time": {
    Type: schema.TypeString,
    Computed: true,
	Description: `Output only. Creation time.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
},
"name": {
    Type: schema.TypeString,
    Computed: true,
	Description: `The generated numeric id for the TagKey.`,
},
"namespaced_name": {
    Type: schema.TypeString,
    Computed: true,
	Description: `Output only. Namespaced name of the TagKey.`,
},
"update_time": {
    Type: schema.TypeString,
    Computed: true,
	Description: `Output only. Update time.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
},
        },
        UseJSONNumber: true,
    }
}



func resourceTagsTagKeyCreate(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
        return err
    }

    obj := make(map[string]interface{})
        parentProp, err := expandTagsTagKeyParent(d.Get( "parent" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
        obj["parent"] = parentProp
    }
        shortNameProp, err := expandTagsTagKeyShortName(d.Get( "short_name" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("short_name"); !isEmptyValue(reflect.ValueOf(shortNameProp)) && (ok || !reflect.DeepEqual(v, shortNameProp)) {
        obj["shortName"] = shortNameProp
    }
        descriptionProp, err := expandTagsTagKeyDescription(d.Get( "description" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
        obj["description"] = descriptionProp
    }


    lockName, err := replaceVars(d, config, "tagKeys/{{parent}}")
    if err != nil {
        return err
    }
    mutexKV.Lock(lockName)
    defer mutexKV.Unlock(lockName)

    url, err := replaceVars(d, config, "{{TagsBasePath}}tagKeys")
    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Creating new TagKey: %#v", obj)
    billingProject := ""



    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
    if err != nil {
        return fmt.Errorf("Error creating TagKey: %s", err)
    }

    // Store the ID now
    id, err := replaceVars(d, config, "tagKeys/{{name}}")
    if err != nil {
        return fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)

    // Use the resource in the operation response to populate
    // identity fields and d.Id() before read
    var opRes map[string]interface{}
    err = tagsOperationWaitTimeWithResponse(
    config, res, &opRes,  "Creating TagKey", userAgent,
        d.Timeout(schema.TimeoutCreate))
    if err != nil {
        // The resource didn't actually create
        d.SetId("")
        return fmt.Errorf("Error waiting to create TagKey: %s", err)
    }


            if err := d.Set("name", flattenTagsTagKeyName(opRes["name"], d, config)); err != nil {
        return err
    }
                                                        
    // This may have caused the ID to update - update it if so.
    id, err = replaceVars(d, config, "tagKeys/{{name}}")
    if err != nil {
        return fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)

    


    log.Printf("[DEBUG] Finished creating TagKey %q: %#v", d.Id(), res)

    return resourceTagsTagKeyRead(d, meta)
}


func resourceTagsTagKeyRead(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
        return err
    }

    url, err := replaceVars(d, config, "{{TagsBasePath}}tagKeys/{{name}}")
    if err != nil {
        return err
    }

    billingProject := ""



    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
    if err != nil {
        return handleNotFoundError(err, d, fmt.Sprintf("TagsTagKey %q", d.Id()))
    }




    if err := d.Set("name", flattenTagsTagKeyName(res["name"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }
    if err := d.Set("parent", flattenTagsTagKeyParent(res["parent"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }
    if err := d.Set("short_name", flattenTagsTagKeyShortName(res["shortName"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }
    if err := d.Set("namespaced_name", flattenTagsTagKeyNamespacedName(res["namespacedName"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }
    if err := d.Set("description", flattenTagsTagKeyDescription(res["description"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }
    if err := d.Set("create_time", flattenTagsTagKeyCreateTime(res["createTime"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }
    if err := d.Set("update_time", flattenTagsTagKeyUpdateTime(res["updateTime"], d, config)); err != nil {
        return fmt.Errorf("Error reading TagKey: %s", err)
    }

    return nil
}

func resourceTagsTagKeyUpdate(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
    	return err
    }

    billingProject := ""



    obj := make(map[string]interface{})
            descriptionProp, err := expandTagsTagKeyDescription(d.Get( "description" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
        obj["description"] = descriptionProp
    }


    lockName, err := replaceVars(d, config, "tagKeys/{{parent}}")
    if err != nil {
        return err
    }
    mutexKV.Lock(lockName)
    defer mutexKV.Unlock(lockName)

    url, err := replaceVars(d, config, "{{TagsBasePath}}tagKeys/{{name}}")
    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Updating TagKey %q: %#v", d.Id(), obj)
updateMask := []string{}

if d.HasChange("description") {
  updateMask = append(updateMask, "description")
}
// updateMask is a URL parameter but not present in the schema, so replaceVars
// won't set it
url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
if err != nil {
  return err
}

    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

    if err != nil {
        return fmt.Errorf("Error updating TagKey %q: %s", d.Id(), err)
    } else {
	log.Printf("[DEBUG] Finished updating TagKey %q: %#v", d.Id(), res)
    }

    err = tagsOperationWaitTime(
        config, res,  "Updating TagKey", userAgent,
        d.Timeout(schema.TimeoutUpdate))

    if err != nil {
        return err
    }

    return resourceTagsTagKeyRead(d, meta)
}

func resourceTagsTagKeyDelete(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
    	return err
    }


    billingProject := ""


    lockName, err := replaceVars(d, config, "tagKeys/{{parent}}")
    if err != nil {
        return err
    }
    mutexKV.Lock(lockName)
    defer mutexKV.Unlock(lockName)

    url, err := replaceVars(d, config, "{{TagsBasePath}}tagKeys/{{name}}")
    if err != nil {
        return err
    }

    var obj map[string]interface{}
    log.Printf("[DEBUG] Deleting TagKey %q", d.Id())

    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
    if err != nil {
        return handleNotFoundError(err, d, "TagKey")
    }

    err = tagsOperationWaitTime(
        config, res,  "Deleting TagKey", userAgent,
        d.Timeout(schema.TimeoutDelete))

    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Finished deleting TagKey %q: %#v", d.Id(), res)
    return nil
}

func resourceTagsTagKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
    config := meta.(*Config)
    if err := parseImportId([]string{
        "tagKeys/(?P<name>[^/]+)",
        "(?P<name>[^/]+)",
    }, d, config); err != nil {
      return nil, err
    }

    // Replace import id for the resource id
    id, err := replaceVars(d, config, "tagKeys/{{name}}")
    if err != nil {
        return nil, fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)


    return []*schema.ResourceData{d}, nil
}

func flattenTagsTagKeyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
    if v == nil {
        return v
    }
	return NameFromSelfLinkStateFunc(v)
}

func flattenTagsTagKeyParent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenTagsTagKeyShortName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenTagsTagKeyNamespacedName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenTagsTagKeyDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenTagsTagKeyCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenTagsTagKeyUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}




func expandTagsTagKeyParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandTagsTagKeyShortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandTagsTagKeyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}
