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



    
func resourceApigeeInstanceAttachment() *schema.Resource {
    return &schema.Resource{
        Create: resourceApigeeInstanceAttachmentCreate,
        Read: resourceApigeeInstanceAttachmentRead,
        Delete: resourceApigeeInstanceAttachmentDelete,

        Importer: &schema.ResourceImporter{
            State: resourceApigeeInstanceAttachmentImport,
        },

        Timeouts: &schema.ResourceTimeout {
            Create: schema.DefaultTimeout(30 * time.Minute),
            Delete: schema.DefaultTimeout(30 * time.Minute),
        },



        Schema: map[string]*schema.Schema{
"environment": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
	Description: `The resource ID of the environment.`,
},
"instance_id": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
	Description: `The Apigee instance associated with the Apigee environment,
in the format 'organisations/{{org_name}}/instances/{{instance_name}}'.`,
},
"name": {
    Type: schema.TypeString,
    Computed: true,
	Description: `The name of the newly created  attachment (output parameter).`,
},
        },
        UseJSONNumber: true,
    }
}



func resourceApigeeInstanceAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
        return err
    }

    obj := make(map[string]interface{})
        environmentProp, err := expandApigeeInstanceAttachmentEnvironment(d.Get( "environment" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("environment"); !isEmptyValue(reflect.ValueOf(environmentProp)) && (ok || !reflect.DeepEqual(v, environmentProp)) {
        obj["environment"] = environmentProp
    }


    lockName, err := replaceVars(d, config, "{{instance_id}}")
    if err != nil {
        return err
    }
    mutexKV.Lock(lockName)
    defer mutexKV.Unlock(lockName)

    url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{instance_id}}/attachments")
    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Creating new InstanceAttachment: %#v", obj)
    billingProject := ""



    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
    if err != nil {
        return fmt.Errorf("Error creating InstanceAttachment: %s", err)
    }

    // Store the ID now
    id, err := replaceVars(d, config, "{{instance_id}}/attachments/{{name}}")
    if err != nil {
        return fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)

    // Use the resource in the operation response to populate
    // identity fields and d.Id() before read
    var opRes map[string]interface{}
    err = apigeeOperationWaitTimeWithResponse(
    config, res, &opRes,  "Creating InstanceAttachment", userAgent,
        d.Timeout(schema.TimeoutCreate))
    if err != nil {
        // The resource didn't actually create
        d.SetId("")
        return fmt.Errorf("Error waiting to create InstanceAttachment: %s", err)
    }


                    if err := d.Set("name", flattenApigeeInstanceAttachmentName(opRes["name"], d, config)); err != nil {
        return err
    }
        
    // This may have caused the ID to update - update it if so.
    id, err = replaceVars(d, config, "{{instance_id}}/attachments/{{name}}")
    if err != nil {
        return fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)

    


    log.Printf("[DEBUG] Finished creating InstanceAttachment %q: %#v", d.Id(), res)

    return resourceApigeeInstanceAttachmentRead(d, meta)
}


func resourceApigeeInstanceAttachmentRead(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
        return err
    }

    url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{instance_id}}/attachments/{{name}}")
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
        return handleNotFoundError(err, d, fmt.Sprintf("ApigeeInstanceAttachment %q", d.Id()))
    }




    if err := d.Set("environment", flattenApigeeInstanceAttachmentEnvironment(res["environment"], d, config)); err != nil {
        return fmt.Errorf("Error reading InstanceAttachment: %s", err)
    }
    if err := d.Set("name", flattenApigeeInstanceAttachmentName(res["name"], d, config)); err != nil {
        return fmt.Errorf("Error reading InstanceAttachment: %s", err)
    }

    return nil
}


func resourceApigeeInstanceAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
    	return err
    }


    billingProject := ""


    lockName, err := replaceVars(d, config, "{{instance_id}}")
    if err != nil {
        return err
    }
    mutexKV.Lock(lockName)
    defer mutexKV.Unlock(lockName)

    url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{instance_id}}/attachments/{{name}}")
    if err != nil {
        return err
    }

    var obj map[string]interface{}
    log.Printf("[DEBUG] Deleting InstanceAttachment %q", d.Id())

    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
    if err != nil {
        return handleNotFoundError(err, d, "InstanceAttachment")
    }

    err = apigeeOperationWaitTime(
        config, res,  "Deleting InstanceAttachment", userAgent,
        d.Timeout(schema.TimeoutDelete))

    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Finished deleting InstanceAttachment %q: %#v", d.Id(), res)
    return nil
}

func resourceApigeeInstanceAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
config := meta.(*Config)

// current import_formats cannot import fields with forward slashes in their value
if err := parseImportId([]string{
		"(?P<instance_id>.+)/attachments/(?P<name>.+)",
		"(?P<instance_id>.+)/(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

// Replace import id for the resource id
id, err := replaceVars(d, config, "{{instance_id}}/attachments/{{name}}")
if err != nil {
	return nil, fmt.Errorf("Error constructing id: %s", err)
}
d.SetId(id)

return []*schema.ResourceData{d}, nil
}

func flattenApigeeInstanceAttachmentEnvironment(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenApigeeInstanceAttachmentName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}




func expandApigeeInstanceAttachmentEnvironment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}
