#include "header.h"

float PermutationAlgo::getMissingRate()
{
    return float(this->missCount) / this->count;
}

int PermutationAlgo::getCount()
{
    return this->count;
}

int PermutationAlgo::getMissCount()
{
    return this->missCount;
}