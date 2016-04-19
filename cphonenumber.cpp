#include <phonenumbers/phonenumberutil.h>
#include "phonenumber.h"

using std::string;
using namespace i18n::phonenumbers;

const PhoneNumberUtil& phone_util(*PhoneNumberUtil::GetInstance());

int isPossibleNumber(char* number, char* region) {
  string numStr(number);
  string regionStr(region);

  return phone_util.IsPossibleNumberForString(numStr, regionStr) ? 1 : 0;
}
