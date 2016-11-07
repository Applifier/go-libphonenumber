#include <phonenumbers/phonenumberutil.h>
#include <cstring>
#include "phonenumber.h"

using std::string;
using namespace i18n::phonenumbers;

const PhoneNumberUtil& phone_util(*PhoneNumberUtil::GetInstance());

int is_possible_number(char* number, char* region) {
  string numStr(number);
  string regionStr(region);

  bool isPossible = phone_util.IsPossibleNumberForString(numStr, regionStr);
  return (isPossible ? 1 : 0);
}

char* get_country_code(char* number) {
  string numStr(number);
  string defaultRegion("ZZ");
  PhoneNumber parsedNumber;

  PhoneNumberUtil::ErrorType error = phone_util.Parse(numStr, defaultRegion, &parsedNumber);
  if (error != PhoneNumberUtil::NO_PARSING_ERROR) {
    return NULL;
  }
  string regionCode;
  phone_util.GetRegionCodeForNumber(parsedNumber, &regionCode);
  return allocAndCopyStr(regionCode.c_str());
}

struct phone_info* parse(char* number, char* region) {
  string numStr(number);
  string regionStr(region);

  PhoneNumber parsedNumber;
  PhoneNumberUtil::ErrorType error = phone_util.Parse(numStr, regionStr, &parsedNumber);
  struct phone_info* res = new_phone_info(number);
  if (error != PhoneNumberUtil::NO_PARSING_ERROR) {
    string error = "Parsing number failed";
    res->error = allocAndCopyStr(error.c_str());
    return res;
  }
  if (!phone_util.IsValidNumber(parsedNumber)) {
    string error = "Invalid number";
    res->error = allocAndCopyStr(error.c_str());
    return res;
  }
  res->valid = 1;
  string formattedNumber;
  phone_util.Format(parsedNumber, PhoneNumberUtil::E164, &formattedNumber);
  res->normalized = allocAndCopyStr(formattedNumber.c_str());
  return res;
}

struct phone_info* new_phone_info(char* number) {
  struct phone_info* pi = (struct phone_info*)malloc(sizeof(struct phone_info));
  pi->valid = 0;
  pi->number = allocAndCopyStr(number);
  pi->normalized = NULL;
  pi->error = NULL;
  return pi;
}

void free_phone_info(struct phone_info* pi) {
  if (pi->number) {
    free(pi->number);
  }
  if (pi->normalized) {
    free(pi->normalized);
  }
  if (pi->error) {
    free(pi->error);
  }
  free(pi);
}

// Helper to copy strings
char* allocAndCopyStr(const char* src) {
  char* dest = (char*)malloc(sizeof(char) * (strlen(src) + 1));
  if (dest == NULL) {
    return NULL;
  }
  strcpy(dest, src);
  return dest;
}
